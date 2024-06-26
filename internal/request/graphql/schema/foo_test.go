package schema

import (
	"context"

	"github.com/sourcenetwork/graphql-go"
	"github.com/sourcenetwork/graphql-go/language/ast"
	"github.com/sourcenetwork/graphql-go/language/printer"

	"github.com/sourcenetwork/defradb/client"
)

func Generate2(definitions []client.CollectionDefinition) (string, error) {
	sm, err := NewSchemaManager()
	if err != nil {
		return "", err
	}
	g := sm.NewGenerator()
	g.Generate(context.TODO(), definitions)

	sdlAst := &ast.Document{}
	for _, t := range g.manager.Schema().TypeMap() {
		if _, ok := t.(*graphql.Object); ok {
			sdlAst.Definitions = append(sdlAst.Definitions, t)
		}
	}

	printer.Print()
}
