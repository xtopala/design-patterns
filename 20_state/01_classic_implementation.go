// State - Classic Implementation

// Before we start talking about State Machines
// as they are really used in the real world, we'll
// take a look at an academic example that is found in
// many books and online examples.

// And the idea of states is being represented by structures
// and having some sort of replacement of one struct with another.

// This will be the demo of the classic implementation
// of a simple state machine with just 2 states.

// So imagine we have a light and the light has
// 2 states, it's on or it's off.
// That's pretty much it.

// But event on the basis of this very simple definition
// what we can do is we can build a rather grotesque looking
// state machine, not the kind we'll find in the real world.
// But once again, this is a classic example, and is one of the
// classic examples.

package main

import "fmt"

// So what we're going to do is we're going
// to have a couple of types.

type Switch struct {
	State State
}

// <- Allows us to turn something On or Off

// In addition to the Switch, just in case someone
// thought this is kind of a self-contained thing,
// we also have this idea of a State.

// Now a state can also be represented as a bunch of
// structs, but we need some sort of interface for the state.

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

// ↑↑↑ Now hold on, this is some really
// 	   heavy over-engineering here.

// Why can't we just keep everything inside a Switch?
// Why do we have to have the state idea?

// We'll see in just a moment, but basically the ide is
// that the switch kind of stays the same but we use implementers
// of this interface and we switch from one implementer to another.

// What we mean by this is that switch has a member called State.
// And when we switch the light on or off we basically replace the
// value of this particular variable, but the way this is done is not
// particularly intuitive shall we say.

// Moving on, what we're going to do is we're going to have a new struct.

type BaseState struct{}

// <- Event though it doesn't contain anything, it is going
//	  to provide the Deep Fold behaviours for the state interface.

// Meaning that these are going to be thing that we
// can subsequently aggregate as part of some other struct.

// So we're going ahead with implementing the methods of
// a State interface.

func (b BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (b BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

// <- This is all weird, but hang on.

// And this is probably very confusing, as to why
// we're defining this Base State which also makes an
// assumption that we haven't really switched the state from
// to another.

// Well, the idea is that we really define
// on and off state subsequently as separate structs.
// Now this is for many reason very inefficient but we can
// define an On State like so.

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

// <- The idea is that whenever we're working with the
//	  OnState the only operation which is allowed on this state
//	  is an operation to turn something off.

// But remember, OnState is what we could say in a way a BaseState,
// meaning it has the same methods, but the which actually turns something
// off here says that the light is already off, so we need to replace it basically.

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning the light off...")
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light is turned off")
	return &OffState{BaseState{}}
}

// And now for the OffState, we do the symmetrical thing.

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning the light on...")
	sw.State = NewOnState()
}

// This is the really complicated way of switching the state.
// So first of all, we have thse On and Off invocataions which
// kind of override what is given by the Base State.
// The Base State provides default implementations.

// Now we need to go back, and do the Switch implementations.

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

// <- What's happening here is effectively double dispatches.

// The king of double dispatche that happens on
// the Visitor design pattern, except that unlike
// the Visitor design pattern what's happening here is
// completly unnecessary.

// Meaning that this entire model can be simplified
// to be much more simpler.

// But let's finish this example.

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

// <- Let's say we're starting with the lights turned off.

// Recap:
// -> Here's the interesting part <- when we try to turn the light off again
// -> The thing about this is that state is already an OffState,
//	  and if we look at it we can see that it doesn't really have an Off method,
//	  it only has an On method
// -> We can invoke it, because we're aggregating the BaseState,
//	  we're actually calling the Off method of BaseState right there
// -> This is purely academic example that we're looking at
// -> Not the kind of stuff that we're likely to be building in
//	  the realn world, but it's a good ilustration
// -> But essentially, the way the state management happens here
//	  is by replacement
// -> And the interesting thing is that the replacement is
//	  done by the state itself

func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
