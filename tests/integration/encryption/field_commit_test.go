// Copyright 2024 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package encryption

import (
	"testing"

	"github.com/sourcenetwork/immutable"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestDocEncryptionField_WithEncryptionOnField_ShouldStoreOnlyFieldsDeltaEncrypted(t *testing.T) {
	test := testUtils.TestCase{
		Actions: []any{
			updateUserCollectionSchema(),
			testUtils.CreateDoc{
				Doc:             john21Doc,
				EncryptedFields: []string{"age"},
			},
			testUtils.Request{
				Request: `
					query {
						commits {
							delta
							docID
							fieldName
						}
					}
				`,
				Results: []map[string]any{
					{
						"delta":     testUtils.NewEncryptedValue(0, 0, immutable.Some("age"), testUtils.CBORValue(21)),
						"docID":     john21DocID,
						"fieldName": "age",
					},
					{
						"delta":     testUtils.CBORValue("John"),
						"docID":     john21DocID,
						"fieldName": "name",
					},
					{
						"delta":     nil,
						"docID":     john21DocID,
						"fieldName": nil,
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}
