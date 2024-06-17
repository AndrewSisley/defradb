// Copyright 2024 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package generator

type Generator interface {
	Generate(maxSize uint64) []byte //
}

type Random struct{}

var _ Generator = (*Random)(nil)

func (g Random) Generate(maxSize uint64) []byte {
	panic("todo")
}
