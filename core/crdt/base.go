// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package crdt

import (
	"context"
	"encoding/binary"

	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"

	"github.com/sourcenetwork/defradb/core"
	"github.com/sourcenetwork/defradb/datastore"
	"github.com/sourcenetwork/defradb/errors"
)

// baseCRDT is embedded as a base layer into all
// the core CRDT implementations to reduce code
// duplication, and better manage the overhead
// tasks that all the CRDTs need to implement anyway
type baseCRDT struct {
	datastore datastore.DSReaderWriter
	headstore datastore.DSReaderWriter
	key       core.DataStoreKey
}

// @TODO paramaterize ns/suffix
func newBaseCRDT(
	datastore datastore.DSReaderWriter,
	headstore datastore.DSReaderWriter,
	key core.DataStoreKey,
) baseCRDT {
	return baseCRDT{
		datastore: datastore,
		headstore: headstore,
		key:       key,
	}
}

func (base baseCRDT) setPriority(
	ctx context.Context,
	key core.DataStoreKey,
	priority uint64,
	cid cid.Cid,
) error {
	prioK := key.ToHeadStoreKey().WithCid(cid)
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, priority)
	if n == 0 {
		return ErrEncodingPriority
	}

	return base.headstore.Put(ctx, prioK.ToDS(), buf[0:n])
}

// get the current priority for given key
func (base baseCRDT) getPriority(ctx context.Context, key core.DataStoreKey, cid cid.Cid) (uint64, error) {
	headStoreKey := key.ToHeadStoreKey().WithCid(cid)
	pbuf, err := base.headstore.Get(ctx, headStoreKey.ToDS())
	if err != nil {
		if errors.Is(err, ds.ErrNotFound) {
			return 0, nil
		}
		return 0, err
	}

	prio, num := binary.Uvarint(pbuf)
	if num <= 0 {
		return 0, ErrDecodingPriority
	}
	return prio, nil
}
