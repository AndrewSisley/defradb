// Copyright 2023 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package descriptions

import (
	"context"
	"encoding/json"

	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/core"
	"github.com/sourcenetwork/defradb/datastore"
)

func SaveCollection(
	ctx context.Context,
	txn datastore.Txn,
	desc client.CollectionDescription,
) error {
	buf, err := json.Marshal(desc)
	if err != nil {
		return err
	}

	collectionKey := core.NewCollectionKey(desc.Name)
	err = txn.Systemstore().Put(ctx, collectionKey.ToDS(), buf)
	if err != nil {
		return err
	}

	collectionSchemaVersionKey := core.NewCollectionSchemaVersionKey(schema.VersionID)
	err = txn.Systemstore().Put(ctx, collectionSchemaVersionKey.ToDS(), buf)
	if err != nil {
		return err
	}

	return nil
}
