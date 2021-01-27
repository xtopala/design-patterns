// State - Handmade State Machine

// In the previous lesson we looked at
// a rather academic example of the classic
// implementation of the state design pattern and
// that's probably not the kind of state machine that we
// actually want to build.

// One of the reasons is that in most cases the states
// and the transitions should not be defined by some heavyweight
// constructs like structures.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// For example we can define states by just defining
// a bunch of constants.

type State int

// Now let's imagine that we want to simulate
// a situation where we're simulating a phone call.
// Pickup a phone and call somebody, and there's lots of
// things that can happen for example we get placed on hold,
// or we get to leave a message or we just get tired of waiting
// and we damn thing down.
// That sort of thing.

// So the idea is that we can represent each of the states
// over the phone with a constant, with an integer in this case.

const (
	OffHook State = iota
	Connecting
	Connected
	OnHold
	OnHook
)

// Now there's one more problem in Go land here,
// and that's the printing of these constants.
// If we would want to print them as strings we would
// have to use a stringer generator or any other kind of generate.

func (s State) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

// <- OK, so these are the states of the system.

// Those are the states the system can be in,
// and now we can have the triggers.

// So the triggers are explicit definitions of
// what can cause us to go from one state to another.
// For example, when we dial a call we can be connected,
// and thus we transition from a state of OffHook to a state Connected.

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

// Now, once again we need to generate a string implementation.

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

// OK, now what we need to do is we need
// to define the rules which transition us from
// one state to another.
// For example when we dial a call, we move from
// the OffHook state to the Connecting state.

// So how can we actually define all these data structures?
// Now unless we're using some external library we can just
// define our own map, which will hold a bunch of different
// trigger results.

type TriggerResult struct {
	Trigger Trigger
	State   State
}

// <- Combination of a trigger and a state we
//	  transition to when that trigger actually happens.

var rules = map[State][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	},
}

// <- It's not just a single trigger result,
//	  because remember from any given state, it might
//	  be possible to transition to more than one state
//	  depending on the trigger.

// Now that we have all of this we can now
// build our state machine and orchestrate this.

// Recap:
// -> This is how we implement a state machine by hand
// -> And this is how we do it in a more realistic setting,
//	  unlike the previous example
// -> All we're doing here is we just have a bunch of constants,
//	  and then we have a map which basically defines all the transition
//	  rules that can happen inside a system

func main() {
	state, exitState := OffHook, OnHook
	// <- when we reach exitState we're done effectively

	for ok := true; ok; ok = state != exitState {
		fmt.Println("The phone is currently:", state)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state][i]
		state = tr.State
	}
	fmt.Println("We're done using the phone")
}
