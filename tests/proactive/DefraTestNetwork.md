# The Defra test network

The Defra test network is a decentralized network of Defra nodes managing, tracking and executing automated tests.

It makes Defra developers day-to-day users of a decentralized, edge-AI, local-first, Defra node network, and could significantly reduce the cost of testing and improve our ability to detect issues.

A happy side-effect of this is the marketing value provided by the system, that not should be factored into the system design, however if seen to be significant perhaps it could be factored into how much effort we wish to commit to the project.

Ideally the solution would be built in such a way that it would be suitable for testing other systems besides Defra, including external projects.

## Overview

### The local dev-test node

A persistent local file based (badger) node will be created locally.  Long term we'll likely want the exact params to be configurable, however initially they can probably be hardcoded.  The local node should be connected to the wider network by default, although that should not be required.

This makes every Defra-developer a Defra-user, running and maintaining a local node connected to the Defra development network.

`make` command(s) should be added to make initializing/refreshing/starting the local server easy, and perhaps the setting up of the test-system configuration.

### Changes to standard test runs

Optionally, when executing the standard test set (e.g. `make test`), a connection to the local db will be attempted, failure to connect should not inhibit the execution of git-committed .go tests.

When a test composed entirely of actions that supports the system is executed by the test framework the results will be added to the local database.  If the test and its actions did not previously exist in the database, they will be added.

Errors when communicating with the local database should not affect the execution of the test run and should be suppressed by default.

Long term, we may want the communication with the database to be async, however given that in most cases the machine will be operating at maximum capacity when running tests anyway, making this async may have a detrimental effect on test run time and would probably be a poor use of development effort.

### Implicit test results

Because the results are tracked against large number of executions across many nodes, some tests may no longer need to explicitly state their expected results and can instead simply assert that the result has not changed, cutting down their development and maintenance costs.

This will not be suitable for all tests, particularly when the test expected results provide valuable documentation, however it could lead to a significant reduction in the cost of our tests.

### Generative tests

The tracking of test failure rates can be fed into the test generator, contributing to it's guessing of which tests to propose and execute and further improving the model.  The generator can propose a mix of existing tests, new tests composed of existing actions, and newly generated actions.

It is hoped that this will actually reduce the time cost of executing tests, as fewer tests may be needed to execute in order to provide the same level of confidence in the code that a traditional test-suite execution would provide.  The level of testing can also be specified - for example a quick run of 1000 tests may be requested locally, before running 10,000 in the CI.

Just as with traditional git-committed tests, generated tests will be logged in the local database and shared with the node network.  It is very important that these can be presented to developers in a readable manner - one long term option might be to export selected tests to compilable .go files, perhaps this could be automatic on test failure.

Initially, it is unlikely that generated tests will be much better than randomly generated ones, and will exist as a nice-to-have supplement to the traditional tests, however, it seems likely that this will gradually improve over time as we invest in improving the generator, and the dataset grows.

### Project risk

This should likely be seen as high-risk (effort) high reward project, and it is unknown how much investment in this system would be required before it begins to yield real gains by developers.  It is very possible that at least initially (if not medium-term) it may even have a negative effect on developer experience due to the overhead in execution cost and system maintenance.

Additionally, while all Defra-devs also becoming day-to-day Defra users is probably a very good thing, both in terms of developer efficacy and marketing, long term it may skew development effort away from paying-user use cases and towards internal use cases.

## Details

### Implicit test result tests

<details>
    <summary>SDL</summary>

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
		codeID: String

		# Test_Action.hash plus codeID
		hash: String @index

		# The hash of the result
		resultHash: String
		# The result accepted as the expected result
		result: Json
	}
</details>

When executing tests with implicit results, the following flow would occur (if system enabled):

1. On test package init:
    1. An http Defra client is created, it will be used by all tests for any test-network operations (the tests themselves still execute using the test-specific clients).
2. Pre test execution:
    1. Check if test contains implicit-result test action, if not, execute test as normal.
    2. Compute the `Test` and `Action` hashes.
3. For each test action:
    1. Execute the action.
    2. Compute the hash of the action's results.
    3. Compute the hash of the `Result` object defined in the SDL - this is essentially the hash of the test hash, the action hash, and the target (git branch/commit/tag/etc).
    4. Fetch the `Result` in the database using the hash computed in the previous step (this could be done in a single query before test execution, but I think fetching one by one the index might actually be more efficient here, and simpler).
    5. Compare the fetched `Result`.`resultHash` with the hash of the actual results.
    6. Optionally, update the `Result` metrics in the database (see [Generative tests](#generative-tests-1)).
    7. Optionally, if comparison fails, write test with difference to file for review by the developer and fail the test

Later, after peer review, the database can be update db using the peer reviewed file.  Tests should not be updated directly from test execution as flaky tests could result in unwanted updates.  Long-term, perhaps we could have a dedicated server for this, and it alone has ACP write rights to protected targets, the server could be an embedded Go client in an app that processes test files (no direct http client access to db).

### Generative tests

<details>
    <summary>SDL</summary>

    type TestSet {
        # optionally, this could contain metadata
	}

	# many-many bridging object
	type Test_TestSet {
		test: Test
		testSet: TestSet
	}

	type Result {
		... # Extending the `Result` type defined in `Implicit test result tests`

		# These are example metrics that may be useful for the generator (they should be added as/when the generator implementation makes use of them).
		# It might be preferable to host these on another, dedicated collection (one-one with Result):

		count: Integer @crdt(type: "pncounter")
		failureCount: Integer @crdt(type: "pncounter")
		# it might be better to make this a pncounter (sum) and then calculate the average (maybe in a view) as the pncounter will handle
		# concurrent updates better than a LWWR, for now, this is just an illustration, so it doesn't really matter.
		averageExecutionTime: Time
	}
</details>

When executing generative tests, the following flow would occur:

1. By owning process (e.g. `make test:generate`)
    1. Generate `TestSet`, `Test`s and `Test_TestSet`s (see SDL)
    2. Write too database any that do not yet exist
2. For each target (`Result`.`codeID`)
    1. Clone target locally (if not current)
    2. Execute command for it to execute the `TestSet` (in a new child process, like the change detector works), writing `Result`s to the database.
3. After the test execution(s) have completed
    1. Get `Result`s for each target and compare, fail test if different
    2. Optionally, write tests with differences to file for review
    3. Optionally, delete target `Result`s

thought: The feature branch `Result`s could live in branch-specific collection with the same schema, this would mean they wont be synced to any subscribing nodes unless opted in to.  The generated tests could also live in another collection, separate from the standard tests for similar P2P reasons.
