// Copyright 2023 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package field_kinds

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestMutationCreateBool_WithValidBool(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Create mutation, boolean field, with a valid boolean",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
				type User {
					name: String
					verified: Boolean
				}`,
			},
			testUtils.CreateDoc{
				Doc: `{
					"name": "John Grisham",
					"verified": true
				}`,
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}
