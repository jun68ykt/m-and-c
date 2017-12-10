//
// An answer by Golang
//       for
// Missionaries and cannibals problem
//
// This problem is described in detail at
// https://en.wikipedia.org/wiki/Missionaries_and_cannibals_problem
//
// written by jun68ykt@github
//
package main

import (
	"fmt"
)

// definition of structure for the river sides
// containing variable 'm' is the number of
// missionaries and 'n' is the number of cannibals
type RiverSide struct {
	m, c int
}

// definition of the state
// which contains 2 river sides, left and right,
// and string variable 'boat' that describes
// which river side the boat
type State struct {
	left, right RiverSide
	boat        string // "left" or "right"
}

// The initial state:
// There are 3 missionaries, 3 cannibals, 1 boat
// in the left side of the river.
// There are none in the right side.
var initialState State = State{
	left:  RiverSide{3, 3},
	right: RiverSide{0, 0},
	boat:  "left",
}

// The goal state:
// There are 3 missionaries, 3 cannibals, 1 boat
// in the right side of the river.
// There are none in the left side.
var goalState State = State{
	left:  RiverSide{0, 0},
	right: RiverSide{3, 3},
	boat:  "right",
}

// Definition of structure for the operation
// which is applied to the state.
// The struct 'Operator' describes the persons
// boarding the boat.
// It contains two int variables, 'm' and 'c'.
// For example Operator {1, 1} is meaning that
// One missionary and one cannibals board the
// boat shipping to opposite shore of the river.
type Operator struct {
	m, c int
}

// The variable 'Operators' is an array of Operator,
// and it contains the selectable options.
// Up to 2 people can ride on our boat, So the var
// 'Operators' is following:
var Operators = [5]Operator{
	{2, 0},
	{1, 0},
	{1, 1},
	{0, 1},
	{0, 2},
}

// Function 'valid' is the validator for the given state.
// It decide whether the given state is safe.
// The rule of this problem says, "On both river sides,
// if the number of cannibals is more than the number of
// missionaries, cannibals eat missionaries.
// So the function 'valid' returns 'true' when
// no missionaries is eaten.
// In addition, this 'valid' checks whether the variables for the
// numbers of people on both sides are not negative.
func valid(state State) bool {
	switch {
	case state.left.m < 0 || state.left.c < 0 || state.right.m < 0 || state.right.c < 0:
		return false
	case state.left.m > 0 && state.left.c > state.left.m:
		return false
	case state.right.m > 0 && state.right.c > state.right.m:
		return false
	default:
		return true
	}
}

// The 'stateTransition' is the state transition
// function of the finite automaton.
// So, it is given one state of current and
// one operator. If the given 'currentState' can
// accept the Operator 'op' and the 'currentState'
// can be properly changed into the next state,
// this function will returns 'nextState'.
// Otherwise, it should report error by 'ok' of false.
func stateTransition(currentState State, op Operator) (nextState State, ok bool) {

	var from, to *RiverSide

	if currentState.boat == "left" {
		from, to = &currentState.left, &currentState.right
		nextState.boat = "right"
		nextState.right = RiverSide{ to.m + op.m, to.c + op.c }
		nextState.left = RiverSide{ from.m - op.m, from.c - op.c }
	} else {
		from, to = &currentState.right, &currentState.left
		nextState.boat = "left"
		nextState.left = RiverSide{ to.m + op.m, to.c + op.c }
		nextState.right = RiverSide{ from.m - op.m, from.c - op.c }
	}

	ok = valid(nextState)

	return
}

func main() {
	// 'solution' is the answer of this problem. It's a slice of Operators and
	// it will be added the Operator element by the func 'solve'.
	solution := []Operator{}

	// 'history' is the map which contains the already visited states
	history := map[State]bool{ initialState: true }

	// the definition of 'solve' as a function value which is a so-called 'closure'
	var solve func(currentState State) bool

	solve = func(currentState State) bool {

		// When the given currentState is the goal, 'solve' returns true.
		if currentState == goalState {
			return true
		}

		// Searching the path to the goal by DFS(Depth-first search)
		for _, op := range Operators {
			if nextState, ok := stateTransition(currentState, op); ok {
				if !history[nextState] {
					history[nextState] = true
					solution = append(solution, op)
					if solve(nextState) {
						return true
					} else {
						solution = solution[0 : len(solution)-1]
					}
				}
			}
		}
		return false
	}

	// Run the solve function given the initial state
	if solve(initialState) {
		fmt.Printf("%+v\n", solution)
	}
}
