// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package test_explain_execute

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestExecuteExplainRequestWithGroup(t *testing.T) {
	test := testUtils.TestCase{

		Description: "Explain (execute) request with groupBy.",

		Actions: []any{
			gqlSchemaExecuteExplain(),

			// Books
			create2AddressDocuments(),

			testUtils.Request{
				Request: `query @explain(type: execute) {
					ContactAddress(groupBy: [country]) {
						country
						_group {
							city
						}
					}
				}`,

				Results: []dataMap{
					{
						"explain": dataMap{
							"executionSuccess": true,
							"sizeOfResult":     1,
							"planExecutions":   uint64(2),
							"selectTopNode": dataMap{
								"groupNode": dataMap{
									"iterations":            uint64(2),
									"groups":                uint64(1),
									"childSelections":       uint64(1),
									"hiddenBeforeOffset":    uint64(0),
									"hiddenAfterLimit":      uint64(0),
									"hiddenChildSelections": uint64(0),
									"selectNode": dataMap{
										"iterations":    uint64(3),
										"filterMatches": uint64(2),
										"scanNode": dataMap{
											"iterations":    uint64(4),
											"docFetches":    uint64(4),
											"filterMatches": uint64(2),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	executeTestCase(t, test)
}