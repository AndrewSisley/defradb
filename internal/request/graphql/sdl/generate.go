// Copyright 2024 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package sdl

import (
	/*
		"github.com/wundergraph/graphql-go-tools/v2/pkg/ast"
		"github.com/wundergraph/graphql-go-tools/v2/pkg/astparser"
		"github.com/wundergraph/graphql-go-tools/v2/pkg/asttransform"
		"github.com/wundergraph/graphql-go-tools/v2/pkg/operationreport"
	*/
	"github.com/sourcenetwork/graphql-go/language/ast"

	"github.com/sourcenetwork/defradb/client"
)

const stdDefs string = `` // todo

func Generate(definitions []client.CollectionDefinition) (string, error) {
	/*
		sdlAst := ast.NewDocument()
		sdlAst.Input.ResetInputBytes([]byte(stdDefs))

		//sdlAst.
		p := astparser.NewParser()
		p.Parse(sdlAst, &operationreport.Report{})

		asttransform.MergeDefinitionWithBaseSchema()
		for _, definition := range definitions {
			defType := ast.ObjectTypeDefinition{
				Description: ast.Description{},
				Name:        ast.ByteSliceReference{},
			}
			sdlAst.AddObjectTypeDefinition(ast.ObjectTypeDefinition{})
		}
	*/

	sdlAst := &ast.Document{
		// +3 for the operation types
		Definitions: make([]ast.Node, len(definitions)+3),
	}

	queryAst := &ast.ObjectDefinition{
		Name: &ast.Name{
			Value: "Query",
		},
	}
	sdlAst.Definitions = append(sdlAst.Definitions, queryAst)
	mutationAst := &ast.ObjectDefinition{
		Name: &ast.Name{
			Value: "Mutation",
		},
	}
	sdlAst.Definitions = append(sdlAst.Definitions, mutationAst)
	subscriptionAst := &ast.ObjectDefinition{
		Name: &ast.Name{
			Value: "Subscription",
		},
	}
	sdlAst.Definitions = append(sdlAst.Definitions, subscriptionAst)

	for _, definition := range definitions {
		fields := definition.GetFields()

		fieldAsts := make([]*ast.FieldDefinition, 0, len(fields))
		for _, field := range fields {
			astType := fieldKindToAstType(field.Kind)
			fieldAst := &ast.FieldDefinition{
				Name: &ast.Name{
					Value: field.Name,
				},
				Type: astType,
			}
			fieldAsts = append(fieldAsts, fieldAst)
		}

		if len(definition.Description.Fields) == 0 {
			// interface def
		} else {
			defAst := &ast.ObjectDefinition{
				Name: &ast.Name{
					Value: definition.GetName(),
				},
				Fields: fieldAsts,
			}
			sdlAst.Definitions = append(sdlAst.Definitions, defAst)
		}
	}

	return astToString(sdlAst), nil
}

func fieldKindToAstType(kind client.FieldKind) ast.Type {
	switch kind.(type) {
	case client.ScalarKind:
		return &ast.Named{
			Name: &ast.Name{
				Value: kind.String(),
			},
		}
	}
	panic("todo")
}

func astToString(sdlAst *ast.Document) string {
	panic("todo")
}
