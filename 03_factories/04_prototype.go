// The Prototype Factory

// Sometimes we'll find ourselves in a situation where we're
// creating lots of very similar objects.

// Like we have a company and we have a couple of fixed positions
// and each of those positions pays a certain fixed amount and we
// want to create a developer struct or employee struct where we have
// the position and the annual income selected from some table or something, I dunno.

// So what we could do here is have what's effectively
// called a -> Prototype Factory
// This is also related to -> Prototype Design Pattern

// But basically we can have pre-configured objects and then we can have
// a factory function which actually operates on and gives us a particular
// pre-configured object.

package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

// <- These are the options that we can create

// Now what we can do is we can create a factory function which is
// going to actually take Developer or Manager as a parameter and depending
// on the value taken it can give us an appropriate employee, which is already
// initialized with some predefined data relative to our company.

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "dev", 175}
	case Manager:
		return &Employee{"", "dead weight", 175000000}
	default:
		panic("unsupported role")
	}
}

// This demonstrates that there is yet another approach to this
// where we have these predefined objects and we have them sort of
// being returned on demand depending on some flag or rather.

// This can work for sort of specifying the kind of object that
// we want created and of course instead of having the actual struct here
// we can have an interface type instead, and that is also going to work just fine.
// Just fine.

func main() {
	m := NewEmployee(Manager)
	m.Name = "Gaffer"
	fmt.Println(m)
}
