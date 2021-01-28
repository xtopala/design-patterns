// State - Switch-Based State Machine

// There is yet another type of state machine
// that needs to be shown, and this state machine
// is very special because instead of having a map of
// the different transitions what happens is we encode that
// information somewhere else.

// And in this particular case, what we do is we
// encoded inside a switch statement.
// So we're going to take a look how that works.

// As usual, first of all let's setup a scenario.
// We're going to try and model a combination lock.

// A combination lock consists of basically 4 digits
// for the lock, and somebody makes up the combination
// and we have to enter the right one to unlock the thing.
// If we enter the wrong combination then the darn thing remains
// locked.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// So again, we're going to have a State as
// a bunch of integers.

type State int

const (
	Locked State = iota
	Failed
	Unlocked
)

// We can model this and we'll notice that
// we're not making any kind of maps here because
// we're going to specify all the transitions from one
// state to another within a switch statement.

// Recap:
// -> This is an example of a different kind of machine in
//	  the sense of the way that it's written so
// -> We do have the states but we don't specify the transitions
//	  explicitly
// -> Instead what happens is we have an infinite loop with
//	  a switch in it and inside that switch we handle every single
// 	  case and we perform the orchestration of the state machine inside
// 	  these cases
// -> To transition from one state to another we simply change some
//	  variable and that's pretty much it
// -> And as the thing runs forever we encounter these variables
// 	  and we simply do a return of something to that effect

func main() {
	code := "1234"
	state := Locked
	entry := strings.Builder{}

	for {
		switch state {
		case Locked:
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
			entry.WriteRune(r)

			// this is the case where the right code has been
			// entered fully, the system is unlocked
			if entry.String() == code {
				state = Unlocked
				break
			}

			// case where we're still entering the digits
			// soon as we enter an icorrect digit, we can detect
			// that fact, and if the code doesn't start with the entry
			// something is wrong
			if strings.Index(code, entry.String()) != 0 {
				state = Failed
			}
		case Failed:
			fmt.Println("FAILED")
			entry.Reset()
			state = Locked
		case Unlocked:
			fmt.Println("UNLOCKED")
			return
		}
	}
}
