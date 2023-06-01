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
	"github.com/lens-vm/lens/host-go/config"
	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/datastore"
	"github.com/sourcenetwork/immutable/enumerable"
)

type JsonDoc = string

type LensRegistry struct {
	lenseWarehouse map[string]*lenseLocker
	// by source v id
	lenseConfigs map[string]client.LensConfig
}

func NewRegistry() *LensRegistry {
	return &LensRegistry{
		lenseWarehouse: map[string]*lenseLocker{},
	}
}

func (r *LensRegistry) RegisterLens(txn datastore.Txn, cfg client.LensConfig) error {
	locker, ok := r.lenseWarehouse[cfg.SourceSchema]
	if !ok {
		locker = newLocker()
		r.lenseWarehouse[cfg.SourceSchema] = locker
	}

	socket := enumerable.NewSocket[JsonDoc]()
	enumerable, err := config.LoadLens[JsonDoc, JsonDoc](r.lenseConfigs[cfg.SourceSchema].Lens, socket)
	if err != nil {
		return err
	}

	locker.returnLense(&lensePipe{
		input:      socket,
		enumerable: enumerable,
	})

	return nil
}

// todo - doc - reset must be called once done!
func (r *LensRegistry) MigrateUp(
	src enumerable.Enumerable[JsonDoc],
	schemaVersionID string,
) (enumerable.Enumerable[JsonDoc], error) {
	lenseLocker, ok := r.lenseWarehouse[schemaVersionID]
	if !ok {
		// todo - error, or return src?
		panic("todo")
	}

	lense := lenseLocker.borrow()
	lense.SetSource(src)

	return lense, nil
}

func (LensRegistry) MigrateDown(
	src enumerable.Enumerable[JsonDoc],
	schemaVersionID string,
) enumerable.Enumerable[JsonDoc] {
	panic("todo")
}

type lensePipe struct {
	input      enumerable.Socket[JsonDoc]
	enumerable enumerable.Enumerable[JsonDoc]
}

type lenseLocker struct {
	// Using a buffered channel provides an easy way to manage a finite
	// number of lenses.
	//
	// We wish to limit this as creating lenses is expensive, and we do not want
	// to be dynamically resizing this collection and spinning up new lense instances
	// in user time, or holding on to large numbers of them.
	safes chan *lensePipe
}

const LENSE_POOL_SIZE int = 5

func newLocker() *lenseLocker {
	return &lenseLocker{
		safes: make(chan *lensePipe, LENSE_POOL_SIZE),
	}
}

func (l *lenseLocker) borrow() enumerable.Socket[JsonDoc] {
	// todo - instead of blocking, if non are available we might want to create temporary instances
	lense := <-l.safes
	return &borrowedEnumerable{
		source: lense,
		locker: l,
	}
}

func (l *lenseLocker) returnLense(lense *lensePipe) {
	// todo - need to reset the wasm state ~here, new lense func required
	l.safes <- lense
}

type borrowedEnumerable struct {
	source *lensePipe
	locker *lenseLocker
}

var _ enumerable.Socket[JsonDoc] = (*borrowedEnumerable)(nil)

func (s *borrowedEnumerable) SetSource(newSource enumerable.Enumerable[JsonDoc]) {
	s.source.input.SetSource(newSource)
}

func (s *borrowedEnumerable) Next() (bool, error) {
	return s.source.enumerable.Next()
}

func (s *borrowedEnumerable) Value() (JsonDoc, error) {
	return s.source.enumerable.Value()
}

func (s *borrowedEnumerable) Reset() {
	s.locker.returnLense(s.source)
	s.source.enumerable.Reset()
}
