// Copyright 2023 Democratized Data Foundation
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

	"github.com/ipfs/go-datastore/query"
	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/core"
	"github.com/sourcenetwork/defradb/datastore"
)

func (db *db) setMigration(ctx context.Context, txn datastore.Txn, cfg client.LensConfig) error {
	// todo - document that source schema version id may not exist locally!
	key := core.NewSchemaVersionMigrationKey(cfg.SourceSchema)

	json, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = txn.Systemstore().Put(ctx, key.ToDS(), json)
	if err != nil {
		return err
	}

	return db.lenseRegistry.RegisterLens(txn, cfg)
}

func (db *db) loadMigrations(ctx context.Context, txn datastore.Txn) error {
	prefix := core.NewSchemaVersionMigrationKey("")
	q, err := txn.Systemstore().Query(ctx, query.Query{
		Prefix: prefix.ToString(),
	})
	if err != nil {
		return err
	}

	for res := range q.Next() {
		// check for Done on context first
		select {
		case <-ctx.Done():
			// we've been cancelled! ;)
			return nil
		default:
			// noop, just continue on the with the for loop
		}

		if res.Error != nil {
			return res.Error
		}

		// now we have a doc key
		//rawDocKey := ds.NewKey(res.Key).BaseNamespace()
		//key := core.NewSchemaVersionMigrationKeyFromString(rawDocKey)

		var cfg client.LensConfig
		err = json.Unmarshal(res.Value, &cfg)
		if err != nil {
			return err
		}

		err = db.lenseRegistry.RegisterLens(txn, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}
