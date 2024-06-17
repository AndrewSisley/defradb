// Copyright 2024 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package proactive

import (
	"context"
	"testing"

	integration "github.com/sourcenetwork/defradb/tests/integration"
	"github.com/sourcenetwork/defradb/tests/proactive/action"
	"github.com/sourcenetwork/defradb/tests/proactive/generator"
)

func CompareBranches(
	ctx context.Context,
	t testing.TB,
	branches []string,
	numberOfTests int,
) {
	generator := &generator.Random{}

	testCases := make([]integration.TestCase, numberOfTests)
	for i := 0; i < numberOfTests; i++ {
		actionTypes := []action.ActionType{
			&action.CreateSchemaActionType{},
		}
		actions, err := Generate(generator, actionTypes...)
		if err != nil {
			t.Fatal(err)
		}

		testActions := make([]any, len(actions))
		for _, action := range actions {
			testActions = append(testActions, action.Executable())
		}

		testCases = append(
			testCases,
			integration.TestCase{
				Actions: testActions,
			},
		)
	}

	for _, branch := range branches {
		// todo - setup branch
		_ = branch

		// how to transfer test cases between branches?
		// Can we use a defra instance for this? - would need to ensure each action is (de)serializable
		// long term, could perhaps this db track the 'performance' of each action and perhaps feed into
		// a neural net/probability based generator (edge ai lol)? (~~need a local option anyway long-term for quick-runs~~
		// this would just be not enabling P2P + subs) - can we push data to the remote auto-magically? (then quick, local runs
		// contribute to the action db and neural-net/etc model(s))
		// could this be supported by P2P, and perhaps Views (allow subscription to subset?)?
		// Could perhaps the actions be strongly typed and not just json? (metadata?)
		//
		// This could be an amazing Defra-use case, turning all defra/source-devs into users and admins of a long-term node network
		// (including schema updates, migrations etc, plus breaking change management).
		// Could this even be decentralized, like git? (should schema/migrations/etc be kept in a repo, allowing easy setup on new
		// machines? Could remain in Defra repo, however maybe a private repo would be better, although.... we could allow external
		// people to contribute and join the network (ACP)! They could even be allowed to contribute their own test actions!)
		// We could include our main integration tests in this set! (warning - performance considerations, fetching test defs from db
		// is slower than compiled code, however we would be able to track which tests are most valuable (to some extent))
		//
		// Can the defra-specifics be abstracted away, allowing solution-useage by other systems?
		//
		// Whilst populating the db from hand-crafted/commited tests would initially be the norm, we could perhaps even work
		// bi-directionally, with the generator adding tests to commitable .go files (problem in where they should be added, naming
		// and other long term complications - perhaps it would be better if trying this, to write them to a 'proposal' file, which devs
		// can either use or ignore)

		for _, testCase := range testCases {
			// todo - we wouldn't actually want to loop through like this (all tests in branch one branch
			// after the other)

			// todo - this will not compare the outputs! Need some magic here
			// - perhaps this process should initiate/own the test db instances?
			// - perhaps `ExecuteTestCase` can be made into an iterator, allowing
			// inspection after each action execution
			integration.ExecuteTestCase(t, testCase)

			// todo - what do we ideally want to happen here?
			//
			// Iterate through each action
			// for each action {
			//   for each branch{
			//     execute and compare - comparision requires data transfer
			//     this could perhaps just be the hash of the results - the actual diff is only useful on failure -
			//     a defra doc id is a hash.
			//
			//     using a defra instance for this would actually open up the possibility of distributing the compute,
			//     possibly using a dedicated machine for common branches (such as develop) to save the need to clone.
			//     The persistance allows each branch to fully complete its run independently - no need to sync - this
			//     in turn means that the action comparison can also be done async, whilst the tests continue.
			//
			//		Using Defra, running the tests async, also means that each execution can be run using the embedded Go
			//		client - the cli/http would be good to bring in later, but they are not a requirement.
			//
			//     We could perhaps use defra to distribute the compute using free github runners, if we can determine
			//     the peer ids before execution.
			//
			//		We can perhaps use channels, to prioritize the local execution of tests over the adding of them to
			//		the db?  In theory this would remove a lot of the performance cost of this system.
			//
			//     It might be possible to eventually apply this technique to the regular change detector. Or, even replace
			//		it, although we'd need to capture more db state than the tests currently do.
			//
			//		Initially, it is not really nessecary to run against mutliple branches auto-magically, the dev could
			//		simply switch branches and then trigger the comparision.  This feature might be useful long-term too
			//		and would allow for easy comparisions of previous executions.  The orchastrating magic could sit on top
			//		of this system.
			//   }
			// }
		}
	}
}

// Reminder - this is likely a high-risk (effort) high reward project, be very mindful of the extra complexity this will
// introduce, and that maintaining it is not free.

const dbSDL string = `
	type Test {
		name: String

		# The test hash is really just a hash of the (ordered by index) hashes of its' actions' hashes, it should not include
		# test name or other metadata
		hash: String @index
	}

	type Action {
		hash: String @index
		json: Json
	}

	# many-many bridging object
	type Test_Action {
		# This action's location within the test, very important for a bunch of reasons,
		# including if the same action is defined multiple times in the same test
		index: Integer

		# Test.hash plus Action.hash plus index (hash just saves space)
		hash: String @index

		test: Test
		action: Action
	}

	type Result {
		test_action: Test_Action # index?
		# The marker used to identify the state of the code, for example the branch name, or git commit id
		codeId: String

		# Test_Action.hash plus codeId
		hash: String @index

		# The hash of the result
		resultHash: String
		# The result accepted as the expected result
		result: Json
	}
`

// simple flow
/*
On test
1. compute action and test hashes

On test action execute
1. hash action result
2. get Result of target using Result.hash (this could be done in a single query, but one by one might actually be more efficient with the hash)
3. compare target result
4. (optional) update Result metrics
5. (optional) write tests with differences to file for review

Later, after peer review
1. Update db using reviewed file - do not write directly from test execution as flaky tests could result in unwanted updates.  We could have
a dedicated server for this, and it alone has ACP write rights to protected branches, the server could be an embedded Go client in an app
that processes test files (no direct http client access to db).
*/

const generativeDbSDL string = `
	type TestSet {

	}

	# many-many bridging object
	type Test_TestSet {
		test: Test
		testSet: TestSet
	}

	type Result {
		...

		# Example metrics that may be useful for the generator (they should be added as/when the generator implementation makes use of them).
		# It might be preferable to host these on another, dedicated collection (one-one with Result):

		count: Integer @crdt(type: "pncounter")
		failureCount: Integer @crdt(type: "pncounter")
		# it might be better to make this a pncounter (sum) and then calculate the average (maybe in a view) as the pncounter will handle
		# concurrent updates better than a LWWR, for now, this is just an illustration, so it doesn't really matter.
		averageExecutionTime: Time
	}
`

// generative flow, comparing branches
/*
By owning process
1. Generate TestSet, Tests and Test_TestSets
2. Write any that do not yet exist

For each codeId
1. Execute TestSet writing Result (TestSet._docID passed as env. var)

After test execution (note, this could be done async using a subscription to Result)
1. Get Results for each codeId and compare, fail test if different
2. (optional) write tests with differences to file for review
3. (optional) delete feature branch Results

thought: The feature branch Results could live in branch-specific collection with the same schema, this would
mean they wont be synced to any subscribing nodes unless opted in to.  The generated tests could also live
in another collection, separate from the standard tests for similar P2P reasons.
*/
