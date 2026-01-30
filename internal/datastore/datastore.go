// Copyright 2025 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package datastore

import (
	"context"

	"github.com/sourcenetwork/corekv"
	"github.com/sourcenetwork/corekv/namespace"
	"github.com/sourcenetwork/defradb/internal/db/lock"
	"github.com/sourcenetwork/defradb/internal/keys"
)

type datastore struct {
	underlying corekv.ReaderWriter

	lockSet *lock.LockSet
}

var _ Keyedstore = (*datastore)(nil)

func newDatastore(rootstore corekv.ReaderWriter, lockSet *lock.LockSet) *datastore {
	return &datastore{
		underlying: namespace.Wrap(rootstore, []byte{dataStoreKey}),
		lockSet:    lockSet,
	}
}

func (s *datastore) Get(ctx context.Context, key Key) ([]byte, error) {
	s.collectionReadLock(ctx, key)

	keyBytes := key.Bytes()
	return s.underlying.Get(ctx, keyBytes)
}

func (s *datastore) Has(ctx context.Context, key Key) (bool, error) {
	s.collectionReadLock(ctx, key)

	keyBytes := key.Bytes()
	return s.underlying.Has(ctx, keyBytes)
}

func (s *datastore) Iterator(ctx context.Context, opts IterOptions) (corekv.Iterator, error) {
	var prefix []byte
	var start []byte
	var end []byte

	if opts.Prefix != nil {
		s.collectionReadLock(ctx, opts.Prefix)
		prefix = opts.Prefix.Bytes()
	}
	if opts.Start != nil {
		s.collectionReadLock(ctx, opts.Start)
		start = opts.Start.Bytes()
	}
	if opts.End != nil {
		end = opts.End.Bytes()
	}

	if opts.Prefix == nil && opts.Start == nil {
		s.lockSet.RLockAll(CtxMustGetTxn(ctx))
	}

	ckvOpts := corekv.IterOptions{
		Prefix:   prefix,
		Start:    start,
		End:      end,
		KeysOnly: opts.KeysOnly,
		Reverse:  opts.Reverse,
	}
	return s.underlying.Iterator(ctx, ckvOpts)
}

func (s *datastore) Set(ctx context.Context, key Key, value []byte) error {
	err := s.collectionWriteLock(ctx, key)
	if err != nil {
		return err
	}

	keyBytes := key.Bytes()
	return s.underlying.Set(ctx, keyBytes, value)
}

func (s *datastore) Delete(ctx context.Context, key Key) error {
	err := s.collectionWriteLock(ctx, key)
	if err != nil {
		return err
	}

	keyBytes := key.Bytes()
	return s.underlying.Delete(ctx, keyBytes)
}

// `Unsafe` returns the underlying `corekv.ReaderWriter`.
//
// This allows access to the underlying the untyped `[]byte` functions and will
// bypass the locksystem.
//
// Legitimate uses of this function are rare, think very carefully if using this.
func (s *datastore) Unsafe() corekv.ReaderWriter {
	return s.underlying
}

func (s *datastore) collectionWriteLock(ctx context.Context, key Key) error {
	colKey, isKeyedByCollection := key.(keys.CollectionedKey)
	if !isKeyedByCollection {
		// No-op, the key does not contain a reference to a collection,
		// so we do not need to lock it
		return nil
	}

	return s.lockSet.CollectionRLockForWrite(CtxMustGetTxn(ctx), colKey.GetCollectionShortID())
}

func (s *datastore) collectionReadLock(ctx context.Context, key Key) {
	colKey, isKeyedByCollection := key.(keys.CollectionedKey)
	if !isKeyedByCollection {
		// No-op, the key does not contain a reference to a collection,
		// so we do not need to lock it
		return
	}

	s.lockSet.CollectionRLockForRead(CtxMustGetTxn(ctx), colKey.GetCollectionShortID())
}
