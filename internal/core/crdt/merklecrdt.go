// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

/*
Package crdt provides CRDT implementations leveraging MerkleClock.
*/
package crdt

import (
	"context"

	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/internal/datastore"
	"github.com/sourcenetwork/defradb/internal/keys"
	"github.com/sourcenetwork/immutable"
)

var FieldCRDTs = []FieldValueCRDT{
	NewLWW(),
	NewCounter(true),
	NewCounter(false),
}

// TryGetFieldCRDT returns the cached instance of the given crdt.
//
// A `NONE_CRDT` will return `nil`.  If nothing is found, `false` will
// be returned, else `true`.
func TryGetFieldCRDT(ct client.CType) (FieldValueCRDT, bool) {
	if ct == client.NONE_CRDT {
		return nil, true
	}

	for _, crdt := range FieldCRDTs {
		if crdt.CType() == ct {
			return crdt, true
		}
	}
	return nil, false
}

type KindLimitedCRDT interface {
	SupportedKinds() []client.FieldKind
}

type Operation struct {
	Name string

	Params map[string]client.FieldKind

	// If true, this operation also has a property of the same name and kind as the field.
	//
	// For example: `set: { name: "John" }`, where `set` is the operation name and `name` is
	// the field name.
	//
	// todo - collision issue?  What if the crdt param names clash with this? (strongly consider
	// defering this problem until implmenting text crdt)
	AcceptsFieldValue bool

	// IncludeAsLegacyGQL constructs the legacy v1 input parameters, and anything supplied by users
	// to them will be used in this operation - for example `input: { name: "John" }`.
	//
	// @deprecated: Remove this property as part of v2.0.0
	IncludeAsLegacyGQL bool
}

type FieldValueCRDT interface {
	CType() client.CType

	String() string

	Description() string

	Operations() []Operation

	Merge(
		ctx context.Context,
		store datastore.Keyedstore,
		key keys.DataStoreKey,
		kind client.FieldKind,
		other Delta,
	) error
}

type DocumentValueCRDT interface {
	Merge(ctx context.Context, store datastore.Keyedstore, key keys.PrimaryDataStoreKey, other Delta) error
}

///

// todo - strongly consider moving this into FieldValueCRDT
type DynamicFieldValueCRDT interface {
	Execute(
		ctx context.Context,
		operation string,
		collectionVersionID string,
		fieldName string,
		value immutable.Option[any],
		priority uint64,
	) (Delta, error)
}
