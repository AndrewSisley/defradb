// Copyright 2026 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package lock

import (
	"sync"

	"github.com/sourcenetwork/corekv"
)

type tryLock struct {
	lock                  *sync.RWMutex
	errorOnCompetingWrite bool
}

func newTryLock(errorOnCompetingWrite bool) *tryLock {
	return &tryLock{
		lock:                  &sync.RWMutex{},
		errorOnCompetingWrite: errorOnCompetingWrite,
	}
}

func (l *tryLock) RLock(isWrite bool) error {
	if isWrite && l.errorOnCompetingWrite {
		nowHoldsLock := l.lock.TryRLock()
		if !nowHoldsLock {
			return corekv.ErrTxnConflict
		}
	} else {
		l.lock.RLock()
	}

	return nil
}

func (l *tryLock) Lock() error {
	if l.errorOnCompetingWrite {
		nowHoldsLock := l.lock.TryLock()
		if !nowHoldsLock {
			return corekv.ErrTxnConflict
		}
	} else {
		l.lock.Lock()
	}

	return nil
}

func (l *tryLock) Unlock() {
	l.lock.Unlock()
}

func (l *tryLock) RUnlock() {
	l.lock.RUnlock()
}
