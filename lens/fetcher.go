// Copyright 2023 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package lens

import (
	"context"
	"encoding/json"

	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/client/request"
	"github.com/sourcenetwork/defradb/core"
	"github.com/sourcenetwork/defradb/datastore"
	"github.com/sourcenetwork/defradb/db/fetcher"
)

// todo - thought - the fetcher stuff would be much nicer as an enumerable I think if not now, open a ticket to change it

type lensedDocumentFetcher struct {
	source fetcher.Fetcher
	lense  *lense
	col    *client.CollectionDescription
}

var _ fetcher.Fetcher = (*lensedDocumentFetcher)(nil)

func (df *lensedDocumentFetcher) Init(
	col *client.CollectionDescription,
	fields []*client.FieldDescription,
	reverse bool,
	showDeleted bool,
) error {
	panic("todo - fetcher sig requires we (re)init all props here")
	return df.source.Init(col, fields, reverse, showDeleted)
}

func (df *lensedDocumentFetcher) Start(ctx context.Context, txn datastore.Txn, spans core.Spans) error {
	return df.source.Start(ctx, txn, spans)
}

func (df *lensedDocumentFetcher) FetchNext(ctx context.Context) (fetcher.EncodedDocument, error) {
	panic("todo")
	return nil, nil
}

func (df *lensedDocumentFetcher) FetchNextDecoded(ctx context.Context) (*client.Document, error) {
	panic("todo - this is used by collection (inclding update)")
	return nil, nil
}

func (df *lensedDocumentFetcher) FetchNextDoc(
	ctx context.Context,
	mapping *core.DocumentMapping,
) ([]byte, core.Doc, error) {
	key, doc, err := df.source.FetchNextDoc(ctx, mapping)
	if err != nil {
		return nil, core.Doc{}, err
	}

	sourceJson, err := coreDocToJson(mapping, doc)
	if err != nil {
		return nil, core.Doc{}, err
	}

	df.lense.Put(doc.SchemaVersionID, sourceJson)

	hasNext, err := df.lense.Next()
	if err != nil {
		return nil, core.Doc{}, err
	}
	if !hasNext {
		// The migration decided to not yield a document, so we cycle through the next fetcher doc
		return df.FetchNextDoc(ctx, mapping)
	}

	migratedDocJson, err := df.lense.Value()
	if err != nil {
		return nil, core.Doc{}, err
	}

	migratedDoc, err := df.jsonToCoreDoc(mapping, migratedDocJson)
	if err != nil {
		return nil, core.Doc{}, err
	}

	return key, migratedDoc, nil
}

func (df *lensedDocumentFetcher) Close() error {
	df.lense.Reset()
	return df.source.Close()
}

func coreDocToJson(mapping *core.DocumentMapping, doc core.Doc) (JsonDoc, error) {
	docAsMap := mapping.ToMap(doc)
	json, err := json.Marshal(docAsMap)
	if err != nil {
		return "", err
	}
	return JsonDoc(json), nil
}

func (df *lensedDocumentFetcher) jsonToCoreDoc(mapping *core.DocumentMapping, jsonDoc JsonDoc) (core.Doc, error) {
	var docAsMap map[string][]byte
	err := json.Unmarshal([]byte(jsonDoc), &docAsMap)
	if err != nil {
		return core.Doc{}, err
	}

	doc := mapping.NewDoc()

	// todo - this is pretty unperformant, change it once tests up and running (loop within a loop - see col.GetField)
	for k, fieldByteValue := range docAsMap {
		if k == request.KeyFieldName {
			doc.SetKey(string(fieldByteValue))
			continue
		}

		fieldDesc, fieldFound := df.col.GetField(k)
		if !fieldFound {
			// todo - but should never happen
		}

		encodedProp := fetcher.encProperty{
			Desc: fieldDesc,
			Raw:  fieldByteValue,
		}
		_, fieldValue, err := encodedProp.Decode()
		if err != nil {
			return core.Doc{}, err
		}
		doc.Fields[fieldDesc.ID] = fieldValue
	}
	doc.SchemaVersionID = df.lense.targetSchemaVersionID

	return doc, nil
}
