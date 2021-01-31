// Template Method - Functiona Template Method

// So now that we've looked at a structural implementation
// of the template method let's take a look at an alternative,
// a functional approach.

// For this, we're not going to have any interfaces or any structs.
// Instead what we're going to do is we're going to make sure that
// the template method operates simply on functions.

// So here's what they can look like.

package main

import "fmt"

// Let's have a function for playing a game,
// kind of similar to the PlayGame function we had
// in the previous demo except this time round we're
// going to take a bunch of arguments.

func PlayGame(
	start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int,
) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}

// <- Different from before, we're unwrapping the interface.

// As a result, we can use this template method inside
// the main function of ours, without really defining any
// interfaces or structs or anything like that.
// We can do everything inside of main.

// So for example if I need a bunch of variables
// for storing the current maximum number of turns,
// or the current player we can have them as members
// right there in the main.

// Recap:
// -> We don't necessarily have to deal with interfaces and structs
// -> Instead, we can define a template method not as something which
//	  uses an interface full of functions,
// -> But actually something which takes functions as arguments and
//    then uses those functions to define the skeleton of some method
// -> Once we have that skeleton, what we do is simply fill in the gaps
//	  so we create every single one of these functions, and then we pass
//	  those functions into the template method use them to actually implement
//    the details of our algorithm

func main() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting a game of chess")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n", turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
