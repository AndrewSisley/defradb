// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package inline_array

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestQueryInlineIntegerArrayWithSumWithFilter(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple inline array, filtered sum of integer array",
		Request: `query {
					users {
						Name
						_sum(FavouriteIntegers: {filter: {_gt: 0}})
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Shahzad",
					"FavouriteIntegers": [-1, 2, -1, 1, 0]
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Shahzad",
				"_sum": int64(3),
			},
		},
	}

	executeTestCase(t, test)
}

func TestQueryInlineNillableIntegerArrayWithSumWithFilter(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple inline array with filter, sum of nillable integer array",
		Request: `query {
					users {
						Name
						_sum(TestScores: {filter: {_gt: 0}})
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Shahzad",
					"TestScores": [-1, 2, null, 1, 0]
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Shahzad",
				"_sum": int64(3),
			},
		},
	}

	executeTestCase(t, test)
}

func TestQueryInlineFloatArrayWithSumWithFilter(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple inline array, filtered sum of float array",
		Request: `query {
					users {
						Name
						_sum(FavouriteFloats: {filter: {_lt: 9}})
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Shahzad",
					"FavouriteFloats": [3.1425, 0.00000000001, 10]
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Shahzad",
				"_sum": 3.14250000001,
			},
		},
	}

	executeTestCase(t, test)
}

func TestQueryInlineNillableFloatArrayWithSumWithFilter(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple inline array with filter, sum of nillable float array",
		Request: `query {
					users {
						Name
						_sum(PageRatings: {filter: {_lt: 9}})
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Shahzad",
					"PageRatings": [3.1425, 0.00000000001, 10, null]
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Shahzad",
				"_sum": float64(3.14250000001),
			},
		},
	}

	executeTestCase(t, test)
}
