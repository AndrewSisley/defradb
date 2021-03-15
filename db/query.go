// Copyright 2020 Source Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.
package db

import (
	"strings"

	"github.com/sourcenetwork/defradb/client"

	gql "github.com/graphql-go/graphql"
)

func (db *DB) ExecQuery(query string) *client.QueryResult {
	res := &client.QueryResult{}
	// check if its Introspection query
	if strings.Contains(query, "IntrospectionQuery") {
		return db.ExecIntrospection(query)
	}

	txn, err := db.NewTxn(false)
	defer txn.Discard()
	if err != nil {
		res.Errors = []interface{}{err.Error()}
		return res
	}

	results, err := db.queryExecutor.ExecQuery(db, txn, query)
	if err != nil {
		res.Errors = []interface{}{err.Error()}
		return res
	}

	if err := txn.Commit(); err != nil {
		res.Errors = []interface{}{err.Error()}
		return res
	}

	res.Data = results
	return res
}

func (db *DB) ExecIntrospection(query string) *client.QueryResult {
	schema := db.schema.Schema()
	// t := schema.Type("userFilterArg")
	// spew.Dump(t.(*gql.InputObject).Fields())
	params := gql.Params{Schema: *schema, RequestString: query}
	r := gql.Do(params)

	res := &client.QueryResult{
		Data:   r.Data,
		Errors: make([]interface{}, len(r.Errors)),
	}

	for i, err := range r.Errors {
		res.Errors[i] = err
	}

	return res
}