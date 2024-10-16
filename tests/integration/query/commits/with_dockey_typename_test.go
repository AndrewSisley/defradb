// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package commits

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestQueryCommitsWithDockeyWithTypeName(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Simple all commits query with dockey and typename",
		Actions: []any{
			updateUserCollectionSchema(),
			testUtils.CreateDoc{
				CollectionID: 0,
				Doc: `{
						"Name":	"John",
						"Age":	21
					}`,
			},
			testUtils.Request{
				Request: `query {
						commits(dockey: "bae-52b9170d-b77a-5887-b877-cbdbb99b009f") {
							cid
							__typename
						}
					}`,
				Results: []map[string]any{
					{
						"cid":        "bafybeihxvx3f7eejvco6zbxsidoeuph6ywpbo33lrqm3picna2aj7pdeiu",
						"__typename": "Commit",
					},
					{
						"cid":        "bafybeih25dvtgei2bryhlz24tbyfdcni5di7akgcx24pezxts27wz7v454",
						"__typename": "Commit",
					},
					{
						"cid":        "bafybeiapquwo7dfow7b7ovwrn3nl4e2cv2g5eoufuzylq54b4o6tatfrny",
						"__typename": "Commit",
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, []string{"users"}, test)
}
