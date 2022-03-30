// The Interface Factory

// When we have a factory function we don't always have to return a struct.
// Instead what we can do is we can return an interface
// that struct conforms to.

package main

import "fmt"

// We'll again stick those Persons, but we're going to do something different.
// Now it will be a private person. *secret secret*

type person struct {
	name string
	age  int
}

// Since this is hidden from the end-user,
// we're going to define an interface, and this
// is going to define the appropriate behaviors for our people here.

type Person interface {
	Fart()
}

func (p *person) Fart() {
	fmt.Printf("Oh hi, my name is %s and I'm %d... THPPTPHTPHPHHPH\n", p.name, p.age)
}

// So this is how a person can actually say something about themselves.
// But when it comes to making a factory function, what we do is we don't
// return a person, we return a -> Interface.

// And the difference here is obviously that now we have just an interface
// to work with, and we can't use it to modify the underlying type.
// Since we're not exposing that person type, only just the interface.

// This is a neat way of encapsulating some information and just having
// our factory expose just the interface that we subsequently work with.
// And that way for example we can have different underlying types, and we can
// imagine that we can have some other structure here.

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) Fart() {
	fmt.Println("well i'm just to tired to... pff")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

// And this becomes interesting because then what we get is different
// underlying object depending on what invocation we've used.

// This way we can have different types of objects in the background.

func main() {
	p1 := NewPerson("Bill", 36)
	p2 := NewPerson("Beowulf", 103)
	p1.Fart()
	p2.Fart()
}
