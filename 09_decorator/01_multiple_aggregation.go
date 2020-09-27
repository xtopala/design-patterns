// The Multiple Aggregation

// One, not so unusual, issue that we all could experience
// is how to combine the functionality of several structures
// inside a single structure.

// With languages that don't support multiple inheritance
// that might become difficult.
// And with Go there is no inheritance.
// There is only aggregation, and there are situations
// where aggregation doesn't give us the desired results.

// We're going to use a classic demo of a Dragon [going medieval]
// which is both a Bird and a Lizard, at the same time.
// So we're going to try to construct a Dragon from
// seperate Bird and Lizard structures.

package main

import "fmt"

type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying!")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling!")
	}
}

// In Go canon ways, what we could do is
// we could simply aggregate both of these.

type Dragon struct {
	Bird
	Lizard
}

// <- Now, this could cause certain problems
// We can see that both Bird and Lizard have an Age field.
// And that's going to be a problem as we try to sort of combine
// the operations of those two.

// Ambiguous selector means we have to specify which age we're referring to.
// The one stored in the Bird or in the Lizard?

// Now the problem with this is that we can introduce a really nasty
// inconsistency into the behaviour of the Dragon if we set the different
// ages to different values.
//  And after all, we don't need separate fields, it's a single age we want
// to keep in a single field.

// One thing that we could do is have a utility method for setting and getting the age.

func (d *Dragon) Age() int {
	return d.Bird.Age
}

func (d *Dragon) SetAge(age int) {
	d.Bird.Age = age
	d.Lizard.Age = age
}

// But it is still very possible to make it inconsistent
// by going into the Dragon explicitly and changing underlying values.
// And that way we can break our operations.
// We could still have inconsistency across the two parts that
// we've aggregated.

// There is no real solution to this in Go.
// There is no language feature that will allow us to
// kind of regularize this whole shenanigans.

// But, we can try to design the entire set of structures
// differently so that this avoided.
// And so that instead of simply aggregating we'll build -> a Decorator

// We'll build proper decorator around the Bird and Lizard types,
// instead od aggregating.

type Aged interface {
	Age() int
	SetAge(age int)
}

// <- This interface defines a contract for having
//	  the Age getters and setters

// Now getters and setters are not particularly idiomatic Go.
// In most cases, we want to avoid having these.
// But, in this unfortunate situation there is no way around it.

// So now we need our types, which will conform to this interface.

type NewBird struct {
	age int
}

func (nb *NewBird) Age() int       { return nb.age }
func (nb *NewBird) SetAge(age int) { nb.age = age }

func (nb *NewBird) Fly() {
	if nb.age >= 10 {
		fmt.Println("Flying!")
	}
}

type NewLizard struct {
	age int
}

func (nl *NewLizard) Age() int       { return nl.age }
func (nl *NewLizard) SetAge(age int) { nl.age = age }

func (nl *NewLizard) Crawl() {
	if nl.age < 10 {
		fmt.Println("Crawling!")
	}
}

// Now, we need a new Dragon.
// And this time around we are going to keep both
// Bird and Lizard inside, but we're not going to do
// the straightforward aggregation, where we automatically
// get all the members.
// Instead, we'll just have fields, one for Bird and one for Lizard.

type BetterDragon struct {
	bird   NewBird
	lizard NewLizard
}

// What we can do next is that we can redefine
// the behaviours such as flying for all and simply redirect them
// or proxy them to the appropriate fields.
// Before we do that, we have to have the getter and the setter for age.

func (d *BetterDragon) Age() int { return d.bird.Age() }
func (d *BetterDragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

// Unlike the previous example, notice that bird and lizard fields
// are lowercased. No direct access here.
// Instead, we can only operate on behaviors.

func (d *BetterDragon) Fly() {
	d.bird.Fly()
}

func (d *BetterDragon) Crawl() {
	d.lizard.Crawl()
}

// So now we have a situation where Dragon needs to be
// initialized by providing instances of Bird nad Lizard
// and this implies that we have to have some sort of factory
// functions for doint that.

func NewBetterDragon() *BetterDragon {
	return &BetterDragon{NewBird{}, NewLizard{}}
}

// Recap:
// -> In the BetterDragon struct we have constructed a Decorator
// -> This constructed object extends the behaviors of the types
//    that we have [Bird and Lizard]
// -> Really what it's doing it's providing better access to the
//    underlying fields of both
// -> And in addition, it combines their the behaviors by providing
//    the interface members with the same names

func main() {
	d := Dragon{}
	// d.Age = 10
	// ↑↑↑ And we can see that this is not allowe: ambiguous selector
	// d.Bird.Age = 10
	// d.Lizard.Age = 10
	// ↑↑↑ Inconsistency, not good
	d.SetAge(5)
	// ↑↑↑ Still not good enough
	d.Bird.Age = 55 // <- We can still do this
	d.Fly()
	d.Crawl()

	// Again!
	bd := BetterDragon{}
	fmt.Printf("Dragon type is: %T\n", bd)
	bd.SetAge(5)
	bd.Fly()
	bd.Crawl()
}
