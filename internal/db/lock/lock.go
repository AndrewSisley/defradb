// Copyright 2025 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package lock

// LockSet manages a set of available locks.
type LockSet struct {
	collectionLockSet *lockSet[uint32]
}

// NewLockSet creates a new LockSet that manages a set of mutexes.
//
// The returned instance is completely independant from any other
// existing LockSet instances.
func NewLockSet() *LockSet {
	return &LockSet{
		collectionLockSet: newLockSet[uint32](),
	}
}

// CollectionLock acquires a write lock for the given collection short id.
//
// This will prevent all other transactions from acquiring a read or write lock
// to the given collection until the lock is released.  The lock will be released
// when the transaction is either committed or discarded.
//
// The acquired lock will not block other threads operating within this transaction.
func (l *LockSet) CollectionLock(txn txn, collectionShortID uint32, errorOnCompetingWrite bool) error {
	// thought: errorOnCompetingWrite (and atomic-ness) can be a function option for all relevant funcs -
	// e.g. users can chose whether a truncate call causes competing calls to error (requires NAC perm?).

	return l.collectionLockSet.Lock(txn, collectionShortID, errorOnCompetingWrite)
}

// CollectionRLock acquires a read lock for the given collection short id.
//
// This will prevent all other transactions from acquiring a write lock
// to the given collection until the lock is released.  The lock will be released
// when the transaction is either committed or discarded.
//
// The read lock can be promoted to a write lock by this transaction, however, currently,
// it does this by first releasing the read lock and then acquiring a write lock.  This can
// permit competing transaction-locks to acquire a write lock, blocking this thread's acquisition
// of the write lock, and allowing both the other transaction's thread, and any previously
// read-locked threads for this transaction to progress concurrently.
//
// If the `isWrite` parameter is set to true, this call may return a `corekv.ErrTxnConflict` if a write lock
// with the `errorOnCompetingWrite` flag set to true is currently held for this collection.  This flag is
// currently set when deleting a collection version, as actions held behind waiting for the release of the
// lock cannot be allowed to continue upon release as the collection may have been deleted.
func (l *LockSet) CollectionRLock(txn txn, collectionShortID uint32, isWrite bool) error {
	return l.collectionLockSet.RLock(txn, collectionShortID, isWrite)
}

func (l *LockSet) CollectionRLockForRead(txn txn, collectionShortID uint32) {
	// Calling `RLock` with `false` will never result in an error, so we can ignore the returned error parameter
	_ = l.CollectionRLock(txn, collectionShortID, false)
}

func (l *LockSet) CollectionRLockForWrite(txn txn, collectionShortID uint32) error {
	return l.CollectionRLock(txn, collectionShortID, true)
}

func (l *LockSet) RLockAll(txn txn) {
	l.collectionLockSet.RLockAll(txn)
}
