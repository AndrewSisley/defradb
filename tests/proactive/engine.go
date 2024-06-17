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
	"github.com/sourcenetwork/defradb/tests/proactive/action"
	"github.com/sourcenetwork/defradb/tests/proactive/generator"
)

func Generate(generator generator.Generator, actionTypes ...action.ActionType) ([]action.Action, error) {
	actions := make([]action.Action, len(actionTypes))

	for i, actionType := range actionTypes {
		maxSize := actionType.Entropy()
		seed := generator.Generate(maxSize)
		action, err := actionType.ConstructValue(seed)
		if err != nil {
			return nil, err
		}

		actions[i] = action
	}

	return actions, nil
}
