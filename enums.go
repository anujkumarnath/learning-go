package main

import (
	"fmt"
)

// NOTE: go doesn't have a separate language feature for enums
// existing language features can used to implement the same

type ServerState int

const (
 // special keyword iota generates successive constant values automatically,
 // in this case 0, 1, 2 and so on
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// implementing fmt.Stringer interface so that the name can be printed
// when the enum int value is passed to fmt.Print*
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	// prints the name, not the integer
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
