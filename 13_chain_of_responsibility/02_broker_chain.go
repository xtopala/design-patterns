// Chain of Responsibility - Broker Chain

// Now, this is a somewhat more advanced
// implementation of the chain of responsibility design pattern.
// It is going to combine multiple approaches, multiple patterns.

// Alogn side of CoR, will also have:
// -> Mediator: by having a centralized component that everyone talks to
// -> Observer: mediator is going to be observable
// -> Command and Query Separations

// We are throwing everything in the kitchen sink into this demo.

// Now, we'll try to replicate the example that we had previuosly.
// The example of the Method Chain, but this time round things are going to be different.
// The creature is going to also have a name, attack and defence values, but the way those
// values are queired is going to be different.

package main

import (
	"fmt"
	"sync"
)

// Lets define a crature.

type Creature struct {
	Name            string
	attack, defense int
	game            *Game
}

// -> In the previous example we only applied the modifiers
// 	  explicitly by callig the Handle method.

// What we wanto to be able to do is, we want to be able to apply
// these modifiers automatically, so as soon we make a modifier passing in
// the creature it automatically gets applied.
// And when we query the creatures attack and defense values we get
// the final calculated value.

// In order to do this the attack and defense values here only store
// the initial values not the calculated values.
// And we're going to grow a whole system of structs on top of this
// in order to make this possible.

// First of all, we're going to introduce a mediator.
// Or, a central component which every creature refers to.
// Now this component is going to be a game, because most creatures
// participate in a game.

// And Game is going to be a centralized component that everyone kind
// of attaches to, or likes talking to, and for that there's lots and lots
// of stuff we need to implement first.

// First, we're going to introduce this idea of Command and Query Separation.
// We're only going to deal with queries and the idea is that
// when you want to get a creature's attack or defense value we make
// a query, which is a separate struct and we send this struct to the creature.
// As opposed to just calling a method on it.

// So how can this work?
// First, we need a type for a Query.

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

// -> Value here it's interesting, apart from the obvious, that that's the value
//    we expect to read, we can also specify the initial value so that
//	  whoever handles this query has a possibility of actually taking an
//	  existing value and modifying it.

// Now we want to define a bunch of interfaces for implementing the Observer.3

type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

// Observable is implemented by whatever type wants to notify other users
// about something happening.

// Now we can finaly build our centralized component, the Game.

type Game struct {
	observers sync.Map
}

// -> This syn map is going to allow us to basically keep a map
//	  of every single subscriber and to iterate this map to go through
// 	  every single subscriber and notify on that subscriber.

// Now that we have this, what we need to be able to do is we need
// to implement the observable interface on the Game.
// Because Game is what every single participlant in the game is
// going to be subscribed to.

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}

		key.(Observer).Handle(q)
		return true
	})
}

// We can now make a constructor for our Crature,
// whisch makes it easier to initilize our little hatchling creature.

func NewCreature(game *Game, name string, attack, defense int) *Creature {
	return &Creature{game: game, Name: name, attack: attack, defense: defense}
}

// The idea is that we don't address attack and defense directly
// instead we have geters.

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)

	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)

	return q.Value
}

// We'll also implement a stringer on a Creature.

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d))", c.Name, c.Attack(), c.Defense())
}

// The question that now faces it self is: How can we implement modifiers?
// Previosly, we had to apply them explicitly, and now they're going
// to be applied implicitly.

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

// -> This is quite simply a template.

// This means that even though we can theoretically give it a Handle method,
// and make it effectively an observer, there's really nothing to put in there.

func (c *CreatureModifier) Handle(q *Query) {
	// nothing !
}

// -> By default, this thing only exists so that we can compose it
//	  as part of actual modifiers.

type DoubleAttackModifier struct {
	CreatureModifier
}

// -> We also want a Handle method, because this needs to be an Observer.

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

// It's not optional for us to have a constructor, we absolutely
// need a Double Attack Modifier constructor.

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{g, c}}
	g.Subscribe(d)

	return d
}

// And the reason here why we need it is before we return our constructed modifier,
// we also need a subscription.
// Whenever wew actually make one, we want that attack modifier to
// participate in the calculation of any values being queried from a creature.

// So whenever the game generates events, the Double Attack Modifier
// gets to process those events, and in our case the queries.

// Now another thing we can do is implement the Closer interface.

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

// The idea here is that we hace a method which can be used to unsubscribe
// this particular modifier from the game events.

// Recap:
// -> This has been much more sophisticated example of how we
//	  would build a mediator with a chain of responsibility on top of it
// -> We can apply a lot of these modifiers one after another, they can
//	  come into the system, they can go out of the system using the Close method
//	  and the state of the Goblin is always going to be *consistent*
// -> Effectively every time we're asking for the goblins attack or defense
//	  we're recalculating it on the basis of a system
// -> This is a more flexible implementation of the Chain of Responsibility

func main() {
	game := &Game{sync.Map{}} // the central mediator
	goblin := NewCreature(game, "Stronk Goblin", 2, 2)
	fmt.Println(goblin.String())

	{ // apply the modifier, but just temporaily
		m := NewDoubleAttackModifier(game, goblin)
		fmt.Println(goblin.String())
		m.Close()
	}

	fmt.Println(goblin.String())
}
