// Copyright 2024 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package action

type ActionType interface {
	Entropy(...any) uint64 // todo - consider returning a larger uint, entropy is probably incorrect name, should probably be []byte
	ConstructValue([]byte) (Action, error)
}

type ActionTypeT[TAction Action] interface { // todo - remove?
	ActionType
	ConstructValueT([]byte) (TAction, error)
}

type Action interface {
	Executable() any
	String() string
}
