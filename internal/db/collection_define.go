// Copyright 2024 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package db

import (
	"context"
	"encoding/json"
	"strings"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/lens-vm/lens/host-go/config/model"
	"github.com/sourcenetwork/immutable"

	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/internal/db/description"
)

func (db *DB) createCollections(
	ctx context.Context,
	newDefinitions []client.CollectionDefinition,
) ([]client.CollectionDefinition, error) {
	returnDescriptions := make([]client.CollectionDefinition, 0, len(newDefinitions))

	existingDefinitions, err := db.getAllActiveDefinitions(ctx)
	if err != nil {
		return nil, err
	}

	newSchemas := make([]client.SchemaDescription, len(newDefinitions))
	for i, def := range newDefinitions {
		newSchemas[i] = def.Schema
	}

	err = setSchemaIDs(newSchemas)
	if err != nil {
		return nil, err
	}

	for i := range newDefinitions {
		newDefinitions[i].Description.SchemaVersionID = newSchemas[i].VersionID
		newDefinitions[i].Schema = newSchemas[i]
	}

	err = db.setCollectionIDs(ctx, newDefinitions)
	if err != nil {
		return nil, err
	}

	err = db.validateNewCollection(
		ctx,
		append(
			append(
				[]client.CollectionDefinition{},
				newDefinitions...,
			),
			existingDefinitions...,
		),
		existingDefinitions,
	)
	if err != nil {
		return nil, err
	}

	txn := mustGetContextTxn(ctx)

	for _, def := range newDefinitions {
		_, err := description.CreateSchemaVersion(ctx, txn, def.Schema)
		if err != nil {
			return nil, err
		}

		if len(def.Description.Fields) == 0 {
			// This is a schema-only definition, we should not create a collection for it
			returnDescriptions = append(returnDescriptions, def)
			continue
		}

		desc, err := description.SaveCollection(ctx, txn, def.Description)
		if err != nil {
			return nil, err
		}

		col := db.newCollection(desc, def.Schema)

		for _, index := range desc.Indexes {
			descWithoutID := client.IndexDescriptionCreateRequest{
				Name:   index.Name,
				Fields: index.Fields,
				Unique: index.Unique,
			}
			if _, err := col.createIndex(ctx, descWithoutID); err != nil {
				return nil, err
			}
		}

		result, err := db.getCollectionByID(ctx, desc.ID)
		if err != nil {
			return nil, err
		}

		returnDescriptions = append(returnDescriptions, result.Definition())
	}

	return returnDescriptions, nil
}

func (db *DB) patchCollection(
	ctx context.Context,
	patchString string,
) error {
	patch, err := jsonpatch.DecodePatch([]byte(patchString))
	if err != nil {
		return err
	}
	existingCols, err := db.getCollections(
		ctx,
		client.CollectionFetchOptions{IncludeInactive: immutable.Some(true)},
	)
	if err != nil {
		return err
	}

	existingColsByID := map[uint32]client.CollectionDescription{}
	existingDefinitions := make([]client.CollectionDefinition, len(existingCols))
	for _, col := range existingCols {
		existingColsByID[col.ID()] = col.Description()
		existingDefinitions = append(existingDefinitions, col.Definition())
	}

	existingDescriptionJson, err := json.Marshal(existingColsByID)
	if err != nil {
		return err
	}

	newDescriptionJson, err := patch.Apply(existingDescriptionJson)
	if err != nil {
		return err
	}

	var newColsByID map[uint32]client.CollectionDescription
	decoder := json.NewDecoder(strings.NewReader(string(newDescriptionJson)))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&newColsByID)
	if err != nil {
		return err
	}
	newDefinitions := make([]client.CollectionDefinition, len(existingCols))
	updatedColsByID := make(map[uint32]struct{})
	for i, col := range existingCols {
		newDefinitions[i].Schema = col.Schema()
		newDefinitions[i].Description = newColsByID[col.ID()]
		updatedColsByID[col.ID()] = struct{}{}
	}
	// append new cols
	for id, col := range newColsByID {
		if _, ok := updatedColsByID[id]; ok {
			continue
		}
		newDefinitions = append(newDefinitions, client.CollectionDefinition{Description: col})
	}

	err = db.validateCollectionChanges(ctx, existingDefinitions, newDefinitions)
	if err != nil {
		return err
	}

	txn := mustGetContextTxn(ctx)
	for _, col := range newColsByID {
		_, err := description.SaveCollection(ctx, txn, col)
		if err != nil {
			return err
		}

		existingCol, ok := existingColsByID[col.ID]
		if ok {
			if existingCol.IsMaterialized && !col.IsMaterialized {
				// If the collection is being de-materialized - delete any cached values.
				// Leaving them around will not break anything, but it would be a waste of
				// storage space.
				err := db.clearViewCache(ctx, client.CollectionDefinition{
					Description: col,
				})
				if err != nil {
					return err
				}
			}

			// Clear any existing migrations in the registry, using this semi-hacky way
			// to avoid adding more functions to a public interface that we wish to remove.

			for _, src := range existingCol.CollectionSources() {
				if src.Transform.HasValue() {
					err = db.LensRegistry().SetMigration(ctx, existingCol.ID, model.Lens{})
					if err != nil {
						return err
					}
				}
			}
			for _, src := range existingCol.QuerySources() {
				if src.Transform.HasValue() {
					err = db.LensRegistry().SetMigration(ctx, existingCol.ID, model.Lens{})
					if err != nil {
						return err
					}
				}
			}
		}

		for _, src := range col.CollectionSources() {
			if src.Transform.HasValue() {
				err = db.LensRegistry().SetMigration(ctx, col.ID, src.Transform.Value())
				if err != nil {
					return err
				}
			}
		}

		for _, src := range col.QuerySources() {
			if src.Transform.HasValue() {
				err = db.LensRegistry().SetMigration(ctx, col.ID, src.Transform.Value())
				if err != nil {
					return err
				}
			}
		}
	}

	return db.loadSchema(ctx)
}

// SetActiveSchemaVersion activates all collection versions with the given schema version, and deactivates all
// those without it (if they share the same schema root).
//
// This will affect all operations interacting with the schema where a schema version is not explicitly
// provided.  This includes GQL queries and Collection operations.
//
// It will return an error if the provided schema version ID does not exist.
func (db *DB) setActiveSchemaVersion(
	ctx context.Context,
	schemaVersionID string,
) error {
	if schemaVersionID == "" {
		return ErrSchemaVersionIDEmpty
	}
	txn := mustGetContextTxn(ctx)

	schema, err := description.GetSchemaVersion(ctx, txn, schemaVersionID)
	if err != nil {
		return err
	}

	colsWithRoot, err := description.GetCollectionsBySchemaRoot(ctx, txn, schema.Root)
	if err != nil {
		return err
	}

	var newName string
	for _, col := range colsWithRoot {
		if col.Name.HasValue() {
			newName = col.Name.Value()
			// Just take the first named, active version
			break
		}
	}

	if newName == "" {
		// If there are no active versions in the collection set, take the name of the schema to be the name of the
		// collection.
		newName = schema.Name
	}

	for _, col := range colsWithRoot {
		col.Name = immutable.Some(newName)
		col.SchemaVersionID = schemaVersionID

		indexDescriptions, err := db.fetchCollectionIndexDescriptions(ctx, col.RootID)
		if err != nil {
			return err
		}
		col.Indexes = indexDescriptions

		col.SchemaVersionID = schema.VersionID

		indexesToRemove := map[int]struct{}{}
		for i, localField := range col.Fields { // todo - must not give new ids to old fields! (make sure this is tested)
			if _, ok := schema.GetFieldByName(localField.Name); !ok {
				indexesToRemove[i] = struct{}{}
			}
		}

		originalFields := col.Fields
		col.Fields = []client.CollectionFieldDescription{}
		for i, field := range originalFields {
			if _, ok := indexesToRemove[i]; !ok {
				col.Fields = append(col.Fields, field)
			}
		}

		for _, globalField := range schema.Fields {
			_, exists := col.GetFieldByName(globalField.Name)
			if !exists {
				col.Fields = append(
					col.Fields,
					client.CollectionFieldDescription{
						Name: globalField.Name,
					},
				)
			}
		}

		err = db.setFieldIDs(ctx, []client.CollectionDefinition{
			{
				Schema:      schema,
				Description: col,
			},
		})
		if err != nil {
			return err
		}

		_, err = description.SaveCollection(ctx, txn, col)
		if err != nil {
			return err
		}
	}

	// Load the schema into the clients (e.g. GQL)
	return db.loadSchema(ctx)
}
