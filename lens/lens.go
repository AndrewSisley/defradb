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
	"github.com/sourcenetwork/immutable/enumerable"
)

type schemaVersionID = string

type lensDoc struct {
	SchemaVersionID schemaVersionID
	JSON            JsonDoc
}

type lense struct {
	lensRegistry *LensRegistry

	lensePipesBySchemaVersionIDs      map[string]enumerable.Concatenation[JsonDoc]
	lenseInputPipesBySchemaVersionIDs map[string]enumerable.Queue[JsonDoc]
	outputPipe                        enumerable.Concatenation[JsonDoc]

	// sorted in asc priority
	schemaVersionHistory []schemaVersionID       // todo - divergance breaks this
	schemaVersionIDIndex map[schemaVersionID]int // cache of schemaVersionHistory index?

	targetSchemaVersionID    schemaVersionID
	targetSchemaVersionIndex int // no point finding this lots of times

	source enumerable.Enumerable[lensDoc]
}

// todo - input via Put, or a source?
var _ enumerable.Enumerable[JsonDoc] = (*lense)(nil)

func New() enumerable.Enumerable[JsonDoc] {
	// todo - remeber that schemaVersionHistory needs too contain versions that do not exist properly locally, and may only exist within
	// migrations.  Thought - docMapper issue with this (I don't think so, as we can only query known versions)?
	return &lense{
		// todo :)
	}
}

// temp - instead of this and a lense-fetcher, we could instead make lense-fetcher (and maybe even fetchers) enumerables
// instead
// todo - thought - this could be interfaced out as an extension of enumerable (within this package)
func (l *lense) Put(schemaVersionID schemaVersionID, value JsonDoc) error {
	panic("todo")
}

func (l *lense) Next() (bool, error) {
	//todo - check output pipe first! Could be items remaining within according to signature

	hasNext, err := l.source.Next()
	if err != nil || !hasNext {
		return false, err
	}

	doc, err := l.source.Value()
	if err != nil {
		return false, err
	}

	if doc.SchemaVersionID == l.targetSchemaVersionID {
		return true, nil
	}

	isMigratingUp := l.schemaVersionIDIndex[doc.SchemaVersionID] < l.targetSchemaVersionIndex

	var inputPipe enumerable.Queue[JsonDoc]
	if p, ok := l.lenseInputPipesBySchemaVersionIDs[doc.SchemaVersionID]; ok {
		inputPipe = p
	} else {
		if isMigratingUp {
			var isWithinHistoryRange bool
			var pipeHead enumerable.Enumerable[JsonDoc]
			for _, schemaVersionID := range l.schemaVersionHistory {
				isDocSchemaVersion := doc.SchemaVersionID == schemaVersionID
				isWithinHistoryRange = isWithinHistoryRange || isDocSchemaVersion

				if !isWithinHistoryRange {
					// If we have not yet reached the fetched schemaVersionID
					// continue until we do so
					continue
				}

				if junctionPipe, ok := l.lensePipesBySchemaVersionIDs[schemaVersionID]; ok {
					// If a pipe already exists for this version, append the pipeHead to the pipeline
					// and break (pipeline is already constructed upstream)
					//
					// pipeHead can safely be assumed to not be nil here
					junctionPipe.Append(pipeHead)
					break
				}

				// The input pipe will be fed documents
				versionInputPipe := enumerable.NewQueue[JsonDoc]()
				l.lenseInputPipesBySchemaVersionIDs[schemaVersionID] = versionInputPipe

				if isDocSchemaVersion {
					inputPipe = versionInputPipe
				}

				// It is a source of the schemaVersion junction pipe, other schema versions
				// may also join as sources to this junction pipe
				junctionPipe := enumerable.Concat[JsonDoc](versionInputPipe)
				l.lensePipesBySchemaVersionIDs[schemaVersionID] = junctionPipe

				// If we have previously laid
				if pipeHead != nil {
					junctionPipe.Append(pipeHead)
				}

				if schemaVersionID == l.targetSchemaVersionID {
					l.outputPipe = junctionPipe
					break
				}

				// The pipe head then becomes the schema version migration to the next version
				// sourcing from any documents at schemaVersionID, or lower schema versions.
				// This also ensures each document only passes through each migration once,
				// in order, and through the same state container (in case migrations use state).
				pipeHead, err = l.lensRegistry.MigrateUp(junctionPipe, schemaVersionID)
				if err != nil {
					return false, err
				}
			}
		} else {
			// todo - reverse
		}
	}

	if inputPipe == nil {
		// no migrations found, return current doc as-is
		return true, nil
	}

	// Place the current doc in the appropriate input pipe
	inputPipe.Put(doc.JSON)

	return l.outputPipe.Next()
}

func (l *lense) Value() (JsonDoc, error) {
	return l.outputPipe.Value()
}

func (l *lense) Reset() {
	// todo - release the lenses!
	panic("todo")
}
