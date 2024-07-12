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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestDocEncryption_WithEncryptionOnLWWCRDT_ShouldStoreCommitsDeltaEncrypted(t *testing.T) {
	test := testUtils.TestCase{
		Actions: []any{
			updateUserCollectionSchema(),
			testUtils.CreateDoc{
				Doc:            john21Doc,
				IsDocEncrypted: true,
			},
			testUtils.Request{
				IgnoreOrder: true,
				Request: `
					query {
						commits {
							collectionID
							delta
							docID
							fieldId
							fieldName
							height
						}
					}
				`,
				Results: []map[string]any{
					{
						"collectionID": int64(1),
						"delta":        testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(21)),
						"docID":        john21DocID,
						"fieldId":      "1",
						"fieldName":    "age",
						"height":       int64(1),
					},
					{
						"collectionID": int64(1),
						"delta":        testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue("John")),
						"docID":        john21DocID,
						"fieldId":      "2",
						"fieldName":    "name",
						"height":       int64(1),
					},
					{
						"collectionID": int64(1),
						"delta":        nil,
						"docID":        john21DocID,
						"fieldId":      "C",
						"fieldName":    nil,
						"height":       int64(1),
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestDocEncryption_UponUpdateOnLWWCRDT_ShouldEncryptCommitDelta(t *testing.T) {
	test := testUtils.TestCase{
		Actions: []any{
			updateUserCollectionSchema(),
			testUtils.CreateDoc{
				Doc:            john21Doc,
				IsDocEncrypted: true,
			},
			testUtils.UpdateDoc{
				Doc: `{
					"age":	22
				}`,
			},
			testUtils.Request{
				Request: `
					query {
						commits(fieldId: "1") {
							delta
						}
					}
				`,
				Results: []map[string]any{
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(22)),
					},
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(21)),
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestDocEncryption_WithMultipleDocsUponUpdate_ShouldEncryptOnlyRelevantDocs(t *testing.T) {
	test := testUtils.TestCase{
		Actions: []any{
			updateUserCollectionSchema(),
			testUtils.CreateDoc{
				Doc:            john21Doc,
				IsDocEncrypted: true,
			},
			testUtils.CreateDoc{
				Doc: islam33Doc,
			},
			testUtils.UpdateDoc{
				DocID: 0,
				Doc: `{
					"age": 22
				}`,
			},
			testUtils.UpdateDoc{
				DocID: 1,
				Doc: `{
					"age": 34
				}`,
			},
			testUtils.Request{
				Request: `
					query {
						commits(fieldId: "1") {
							delta
							docID
						}
					}
				`,
				Results: []map[string]any{
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(22)),
						"docID": john21DocID,
					},
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(21)),
						"docID": john21DocID,
					},
					{
						"delta": testUtils.CBORValue(34),
						"docID": islam33DocID,
					},
					{
						"delta": testUtils.CBORValue(33),
						"docID": islam33DocID,
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestDocEncryption_WithEncryptionOnCounterCRDT_ShouldStoreCommitsDeltaEncrypted(t *testing.T) {
	const docID = "bae-d3cc98b4-38d5-5c50-85a3-d3045d44094e"

	test := testUtils.TestCase{
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
                    type Users {
                        points: Int @crdt(type: "pcounter")
                    }
                `},
			testUtils.CreateDoc{
				Doc:            `{ "points": 5 }`,
				IsDocEncrypted: true,
			},
			testUtils.Request{
				Request: `
					query {
						commits {
							delta
							docID
						}
					}
				`,
				Results: []map[string]any{
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(5)),
						"docID": docID,
					},
					{
						"delta": nil,
						"docID": docID,
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestDocEncryption_UponUpdateOnCounterCRDT_ShouldEncryptedCommitDelta(t *testing.T) {
	test := testUtils.TestCase{
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
                    type Users {
                        points: Int @crdt(type: "pcounter")
                    }
                `},
			testUtils.CreateDoc{
				Doc:            `{ "points": 5 }`,
				IsDocEncrypted: true,
			},
			testUtils.UpdateDoc{
				Doc: `{
					"points": 3
				}`,
			},
			testUtils.Request{
				Request: `
					query {
						commits(fieldId: "1") {
							delta
						}
					}
				`,
				Results: []map[string]any{
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(3)),
					},
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(5)),
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestDocEncryption_UponEncryptionSeveralDocs_ShouldStoreAllCommitsDeltaEncrypted(t *testing.T) {
	test := testUtils.TestCase{
		Actions: []any{
			updateUserCollectionSchema(),
			testUtils.CreateDoc{
				Doc:            "[" + john21Doc + ", " + islam33Doc + "]",
				IsDocEncrypted: true,
			},
			testUtils.Request{
				Request: `
					query {
						commits {
							delta
							docID
						}
					}
				`,
				Results: []map[string]any{
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue(21)),
						"docID": testUtils.NewDocIndex(0, 0),
					},
					{
						"delta": testUtils.NewEncryptedValue(0, 0, immutable.None[string](), testUtils.CBORValue("John")),
						"docID": testUtils.NewDocIndex(0, 0),
					},
					{
						"delta": nil,
						"docID": testUtils.NewDocIndex(0, 0),
					},
					{
						"delta": testUtils.NewEncryptedValue(0, 1, immutable.None[string](), testUtils.CBORValue(33)),
						"docID": testUtils.NewDocIndex(0, 1),
					},
					{
						"delta": testUtils.NewEncryptedValue(0, 1, immutable.None[string](), testUtils.CBORValue("Islam")),
						"docID": testUtils.NewDocIndex(0, 1),
					},
					{
						"delta": nil,
						"docID": testUtils.NewDocIndex(0, 1),
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestDocEncryption_IfTwoDocsHaveSameFieldValue_CipherTextShouldBeDifferent(t *testing.T) {
	test := testUtils.TestCase{
		Actions: []any{
			updateUserCollectionSchema(),
			testUtils.CreateDoc{
				Doc: `{
						"name": "John",
						"age": 21
					}`,
				IsDocEncrypted: true,
			},
			testUtils.CreateDoc{
				Doc: `{
						"name": "Islam",
						"age": 21
					}`,
				IsDocEncrypted: true,
			},
			testUtils.Request{
				Request: `
					query {
						commits(fieldId: "1") {
							delta
							fieldName
						}
					}
				`,
				Asserter: testUtils.ResultAsserterFunc(func(_ testing.TB, result []map[string]any) (bool, string) {
					require.Equal(t, 2, len(result), "Expected 2 commits")
					require.Equal(t, result[0]["fieldName"], "age")
					delta1 := result[0]["delta"]
					delta2 := result[1]["delta"]
					assert.NotEqual(t, delta1, delta2, "docs should be encrypted with different encryption keys")
					return true, ""
				}),
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}
