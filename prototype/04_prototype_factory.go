// The Prototype Factory

// Alright so now that we've figured out how to get
// object copying to work the way that we expect it to work,
// how about an example where we setup a system where it's easier
// to actually use these prototypes.

// This time around we'll have employee and address.
// And the idea is when we have employees they might work
// in different offices of one company.

// Essentially the problem is that we have too much customization
// being done by hand and it would not be nice to take this customization
// and put this into some sort of a set of functions for example.

// And this is what we typically call a -> Prototype Factory.

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

// Here is really immaterial whether the address is a pointer or
// a value, both of these are going to work.

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek inside
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)

	return &result
}

// Really the scenario is going to be slightly different so what
// we want to be able to do is we want to be able to have a couple
// of predefined addresses.

// So we can have a variable called the main office and this will
// be sort of half configured Employee struct.

var mainOffice = Employee{"", Address{0, "123 East Dr", "London"}}
var auxOffice = Employee{"", Address{0, "66 West Dr", "London"}}

// Now we want some sort of utility function for making a new
// employee based on that particular prototype.
// Because what these things are prototypes, we want to make a
// copy of the prototype and then allow the person to customize it
// that's how our prototype factories typically do it.

func newEmployee(proto *Employee, name string, suite int) *Employee {
	res := proto.DeepCopy()
	res.Name = name
	res.Office.Suite = suite

	return res
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

// So what is the point of this sort of factory approach?
// The point is we make it more convenient to actually create
// these objects.

// So instead of making copies by hand, what we do is we make
// the copies, we perform customization using a function as opossed
// to just statements we write one after another.

// This Prototype Factory approach is simply a convenience approach
// which makes the Prototype Factory a lot easier to use.

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
