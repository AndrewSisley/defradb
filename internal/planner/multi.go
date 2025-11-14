// Copyright 2024 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package planner

import (
	"fmt"

	"github.com/sourcenetwork/defradb/internal/core"
	"github.com/sourcenetwork/defradb/internal/keys"
)

/*
A MultiNode is a planNode which contains multiple sub nodes,
that can be executed either in parallel, and serial. Each Values()
response is added to the stored document. Each child node is a named
planNode, where the name is the target field for the planNode.

This is also the basis of the MultiScannerNode. The MultiScannerNode
is a MultiNode, which shares an underlying scanNode. Each step of a
MultiScannerNode takes one value from the source node, and uses its
results in all the attached multinodes.
*/

type MultiNode interface {
	planNode
	Children() []planNode
}

// parallelNode implements the MultiNode interface. It
// enables parallel execution of planNodes. This is needed
// if a single request has multiple Select statements at the
// same depth in the request.
// Eg:
//
//	user {
//			_docID
//			name
//			friends {
//				name
//			}
//			_version {
//				cid
//			}
//	}
//
// In this example, both the friends selection and the _version
// selection require their own planNode sub graphs to complete.
// However, they are entirely independent graphs, so they can
// be executed in parallel.
type parallelNode struct { // serialNode?
	documentIterator
	docMapper

	p *Planner

	children     []planNode
	childIndexes []int

	source planNode
}

func (p *parallelNode) applyToPlans(fn func(n planNode) error) error {
	for _, plan := range p.children {
		if err := fn(plan); err != nil {
			return err
		}
	}
	return nil
}

func (p *parallelNode) Kind() string {
	return "parallelNode"
}

func (p *parallelNode) Init() error {
	newChildren := make([]planNode, len(p.children))
	newChildIndexes := make([]int, len(p.childIndexes))
	/*
		endIndex := len(p.children) - 1
		startIndex := 0
		for i, child := range p.children {
			switch child.(type) {
			case *dagScanNode:
				newChildren[endIndex] = child
				newChildIndexes[endIndex] = p.childIndexes[i]
				endIndex--
			default:
				newChildren[startIndex] = child
				newChildIndexes[startIndex] = p.childIndexes[i]
				startIndex++
			}
		}
	*/
	startIndex := 0
	for i, child := range p.children {
		switch child.(type) {
		case *scanNode:
			newChildren[startIndex] = child
			newChildIndexes[startIndex] = p.childIndexes[i]
			startIndex++
		}
	}
	for i, child := range p.children {
		switch child.(type) {
		case *dagScanNode, *scanNode:
			// noop
		default:
			newChildren[startIndex] = child
			newChildIndexes[startIndex] = p.childIndexes[i]
			startIndex++
		}
	}
	for i, child := range p.children {
		switch child.(type) {
		case *dagScanNode:
			newChildren[startIndex] = child
			newChildIndexes[startIndex] = p.childIndexes[i]
			startIndex++
		}
	}
	p.children = newChildren
	p.childIndexes = newChildIndexes

	return p.applyToPlans(func(n planNode) error {
		return n.Init()
	})
}

func (p *parallelNode) Start() error {
	return p.applyToPlans(func(n planNode) error {
		return n.Start()
	})
}

func (p *parallelNode) Prefixes(prefixes []keys.Walkable) {
	_ = p.applyToPlans(func(n planNode) error {
		n.Prefixes(prefixes)
		return nil
	})
}

func (p *parallelNode) Close() error {
	return p.applyToPlans(func(n planNode) error {
		return n.Close()
	})
}

// Next loops through all the children nodes, and calls Next().
// It only needs a single child plan to return true for it
// to return true. Same with errors.
func (p *parallelNode) Next() (bool, error) {
	p.currentValue = p.documentMapping.NewDoc()
	println("pn next------")
	println(len(p.children))
	var orNext bool
	for i, plan := range p.children {
		println(fmt.Sprintf("%T", plan))
		var next bool
		var err error
		// isMerge := false
		switch n := plan.(type) {
		case *scanNode, *typeIndexJoin:
			// isMerge = true
			next, err = p.nextMerge(i, n)
		case *dagScanNode:
			next, err = p.nextAppend(i, n)
		default:
			panic(fmt.Sprintf("%T", n))
		}
		if err != nil {
			return false, err
		}
		println(next)
		orNext = orNext || next
	}
	// if none of the children return true for next, then this will be false.
	// if ANY of the children return true, this will be true (logical OR)
	return orNext, nil
}

func (p *parallelNode) nextMerge(_ int, plan planNode) (bool, error) {
	if next, err := plan.Next(); !next {
		return false, err
	}

	// Field-by-fields check is necessary because parallelNode can have multiple children, and
	// each child can return the same doc, but with different related fields available
	// depending on what is requested.
	newFields := plan.Value().Fields
	for i := range newFields {
		if p.currentValue.Fields[i] == nil {
			p.currentValue.Fields[i] = newFields[i]
		}
	}

	return true, nil
}

func (p *parallelNode) nextAppend(index int, plan planNode) (bool, error) {
	key := p.currentValue.GetID()
	if key == "" {
		return false, nil
	}

	// pass the doc key as a reference through the prefixes interface
	prefixes := []keys.Walkable{keys.DataStoreKey{DocID: key}}
	plan.Prefixes(prefixes)
	err := plan.Init()
	if err != nil {
		return false, err
	}

	results := make([]core.Doc, 0)
	for {
		next, err := plan.Next()
		if err != nil {
			return false, err
		}

		if !next {
			break
		}

		results = append(results, plan.Value())
	}
	p.currentValue.Fields[p.childIndexes[index]] = results
	return true, nil
}

func (p *parallelNode) Source() planNode { return p.source }

func (p *parallelNode) Children() []planNode {
	return p.children
}

func (p *parallelNode) addChild(fieldIndex int, node planNode) {
	p.children = append(p.children, node)
	p.childIndexes = append(p.childIndexes, fieldIndex)
}

func (n *selectNode) addSubPlan(fieldIndex int, newPlan planNode) error {
	println("addSubPlan")
	println(fieldIndex)
	println(fmt.Sprintf("%T", newPlan))
	println(fmt.Sprintf("%T", n.source))
	switch sourceNode := n.source.(type) {
	case *scanNode, *pipeNode, *typeIndexJoin:
		parallelNode := &parallelNode{
			p:         n.planner,
			source:    newPlan,
			docMapper: docMapper{n.source.DocumentMap()},
		}
		_, newIsJoin := newPlan.(*typeIndexJoin)
		_, srcIsScan := n.source.(*scanNode)
		if newIsJoin && srcIsScan {
			//panic("fdas")
			//panic(fmt.Sprintf("%T", newPlan))
			//if _, ok := newPlan.(*scanNode); !ok {
			if _, ok := n.source.(*scanNode); !ok {
				//parallelNode.addChild(-1, n.source)
			}
			//}
		} else {
			//parallelNode.addChild(-1, n.source)
		}

		parallelNode.addChild(-1, n.source)
		parallelNode.addChild(fieldIndex, newPlan)
		n.source = parallelNode

	// we already have an existing parallelNode as our source
	case *parallelNode:
		if _, ok := newPlan.(*typeIndexJoin); ok {
			for i, child := range sourceNode.children {
				if _, ok := child.(*scanNode); ok {
					sourceNode.children[i] = newPlan
					sourceNode.childIndexes[i] = fieldIndex
					return nil
				}
			}
		}

		sourceNode.addChild(fieldIndex, newPlan)

	default:
		panic(fmt.Sprintf("%T", n.source))
	}
	return nil
}
