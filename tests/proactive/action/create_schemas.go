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

import (
	"bytes"
	"math"
	"strings"

	"github.com/icza/bitio"

	"github.com/sourcenetwork/defradb/client"
	integration "github.com/sourcenetwork/defradb/tests/integration"
)

type CreateSchemaActionType struct {
	numberOfSchemasToGenerate  uint64
	maxNumberOfFieldsPerSchema float64
	numberOfScalarFieldTypes   int
}

var _ ActionType = (*CreateSchemaActionType)(nil)

type CreateSchema struct {
	integration.SchemaUpdate
}

var _ Action = (*CreateSchema)(nil)

func Configure(...any) {

}

func (a *CreateSchemaActionType) Entropy(...any) uint64 {
	var numberOfSchemasToGenerate uint64 = 1   // todo - should/can this be 'max'?
	var maxNumberOfFieldsPerSchema float64 = 5 // todo - make these params
	numberOfScalarFieldTypes := len(client.FieldKindStringToEnumMapping)

	fieldTypes := numberOfScalarFieldTypes + (2 * int(numberOfSchemasToGenerate))

	entropy := math.Pow(maxNumberOfFieldsPerSchema, float64(fieldTypes)) // todo - can this just be '*'?
	return uint64(entropy)
}

func (a *CreateSchemaActionType) ConstructValue(seed []byte) (Action, error) { // todo - 'seed' is probably incorrect name
	fieldTypes := a.numberOfScalarFieldTypes + (2 * int(a.numberOfSchemasToGenerate)) // todo - cache
	sizeOfEachSchema := int(a.maxNumberOfFieldsPerSchema) * fieldTypes

	byteBuf := bytes.NewBuffer(seed)
	buf := bitio.NewReader(byteBuf)

	bitsPerFieldF := math.Log2(float64(fieldTypes))
	bitsPerField := roundUp(bitsPerFieldF)

	collections := make([][]client.FieldKind, 0, int(a.numberOfSchemasToGenerate))
	for collectionIndex := 0; collectionIndex < int(a.numberOfSchemasToGenerate); collectionIndex++ {
		fieldKinds := make([]client.FieldKind, 0, int(a.maxNumberOfFieldsPerSchema))

		for i := 0; i < sizeOfEachSchema; i++ {
			fieldBits, err := buf.ReadBits(uint8(bitsPerField))
			if err != nil {
				return nil, err
			}
			fieldKind := a.bitsToFieldKind(fieldBits)
			fieldKinds = append(fieldKinds, fieldKind)
		}

		collections = append(collections, fieldKinds)
	}

	sdlBuilder := strings.Builder{}
	for _, collection := range collections {
		panic("todo")
		for _, _ = range collection {

		}
	}

	sdl := sdlBuilder.String()
	return &CreateSchema{
		integration.SchemaUpdate{
			Schema: sdl,
		},
	}, nil
}

func (a *CreateSchema) String() string {
	panic("todo")
}

func (a *CreateSchema) Executable() any {
	return a.SchemaUpdate
}

func (a *CreateSchemaActionType) bitsToFieldKind(bits uint64) client.FieldKind {
	switch bits {
	case 0:
		return client.FieldKind_DocID
	case 1:
		return client.FieldKind_NILLABLE_BOOL
	// etc
	default:
		minusScalars := bits - uint64(len(client.FieldKindStringToEnumMapping))
		isID := math.Remainder(float64(minusScalars), 2) > 0
		if isID {
			return client.FieldKind_DocID
		}
		objectID := int(minusScalars / 2)
		return client.ObjectKind(objectID) // todo - '2' is incorrect everywhere, can be one-many :)
	}
}

func roundUp(f float64) uint64 { // todo - this is very lazy and probably incorrect (def for negative numbers)
	r := uint64(f)
	if f-float64(r) > 0 {
		return r + 1
	}
	return r
}
