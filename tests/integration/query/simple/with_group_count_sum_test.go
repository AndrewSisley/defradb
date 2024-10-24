// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package simple

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestQuerySimpleWithGroupByStringWithInnerGroupBooleanAndSumOfCount(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with group by string, with child group by boolean, and sum of count",
		Request: `query {
					users(groupBy: [Name]) {
						Name
						_sum(_group: {field: _count})
						_group (groupBy: [Verified]){
							Verified
							_count(_group: {})
						}
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John",
					"Age": 25,
					"Verified": true
				}`,
				`{
					"Name": "John",
					"Age": 32,
					"Verified": true
				}`,
				`{
					"Name": "John",
					"Age": 34,
					"Verified": false
				}`,
				`{
					"Name": "Carlo",
					"Age": 55,
					"Verified": true
				}`,
				`{
					"Name": "Alice",
					"Age": 19,
					"Verified": false
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "John",
				"_sum": int64(3),
				"_group": []map[string]any{
					{
						"Verified": true,
						"_count":   int(2),
					},
					{
						"Verified": false,
						"_count":   int(1),
					},
				},
			},
			{
				"Name": "Alice",
				"_sum": int64(1),
				"_group": []map[string]any{
					{
						"Verified": false,
						"_count":   int(1),
					},
				},
			},
			{
				"Name": "Carlo",
				"_sum": int64(1),
				"_group": []map[string]any{
					{
						"Verified": true,
						"_count":   int(1),
					},
				},
			},
		},
	}

	executeTestCase(t, test)
}
