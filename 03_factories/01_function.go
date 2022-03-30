// The Factory Function [a.k.a. Constuctor]

// In most cases with Go, we can take a struct and
// we can just initialize that struct using the {} notation.
// But, here's more buts!

package main

import "fmt"

// So if we, again, make a Person struct, like so:

type Person struct {
	Name string
	Age  int
}

// However there are situations when we want to have some sort of
// behaviour happen as an object is created, as that struct is created.

// We want default values, for example, and let's say there's another
// property here, another field, which counts how many eyes a person has. Yes.

// In most cases, they all have 2 eyes, in 99.999% of recurring cases we wouldn't
// really want to customize this field. Growse !

type PersonWithEyes struct {
	Name     string
	Age      int
	EyeCount int
}

// So we can do this with some sort of -> Factory Function
// And this function is nothing more than a freestanding function
// which returns an instance of this struct that we want to create.

func NewPerson(name string, age int) *PersonWithEyes {
	return &PersonWithEyes{name, age, 2} // 2 eyes is a good default
}

// Now if that wandering wizard ever gets in a fight,
// we can always modify his line of sight.

// Recap:
// -> We see factory functions all over the place when there's some logic
//    that needs to be applied when initializing a particular object
// -> Simple objects we can initialize them with just {}
// -> But if we want additional stuff to happen as we're creating a particular
//    struct that's when we could use these factory functions

// We also could check that when we are constructing that some logic is actually valid.

func NewWizardPerson(name string, age int) *Person {
	if age < 16 {
		panic("this is not a legal age for wizardry")
	}

	return &Person{name, age}
}

func main() {
	s := Person{"Saruman", 25}
	g := NewPerson("Gandalf", 22)
	g.EyeCount = 1

	fmt.Println(s)
}
