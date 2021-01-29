// Template Method

// Imagine that we're simulating different kinds of games.
// Now these games all have very similar structure.

// So if we think about games like chess or checkers
// or a card game, then they're all pretty much the same
// in the sense that you have a bunch of players and each of
// the players takes their turn.
// And there are of course variations.

// But let's imagine a simple scenario, like a game of
// chess or checkers for example, where we simply have two
// players, and every signle player takes a turn one after another.

// What we can do is we can formalize this process using
// a -> Template Method.

package main

import "fmt"

// But in order to do this, we also need some sort of
// interface for what parts of the game we're interested in.

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

// Now that we have this interface, we can write
// a template method which is simply a skeleton algorithm
// which makes use it.

func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

// <- So the idea here is that we simply use the interface
//	  and invoke the interface members in the exact order in
//	  which we want to define the algorithm.

// Now that we have this, we can make the actual game,
// a game of chess for example.
// Adn of course, here we'll not going to implement all
// of the chess rules and whatever, we just want to show a
// simulation of how the game can proceed under this template
// method paradigm.

// Lets have a type for Chess.

type chess struct {
	turn, maxTurns, currentPlayer int
}

// Since chess is a game, it must implement Game interface.

func (c *chess) Start() {
	fmt.Println("Starting a new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer // alternate between 0 and 1, forever
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

// <- This factory returns Game interface,
//	  and that way we can sort of hide all the internals
//	  of the chess struct we've created.

// Recap:
// -> Essentially the template method is a skeleton algorithm
// -> We can see that with PlayGame() method we're using the abstract
//	  members in a way
// -> Those are the members of an interface which we don't have a
//	  concrete implementation
// -> Until we actually implement the Game interface in some struct of
//    ours and then we pass that struct into the play game function and we
//	  actually use it
// -> We actually use that high level algorithm with the definitions of
//	  the different members being defined in whoever actually implemented
//	  this interface

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}
