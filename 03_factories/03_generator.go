// The Factory Generator

// Now we're going to take a look at the idea of
// generating factories.
// It's somewhat different approach to what we've seen before.

package main

import "fmt"

// We're going to work here with some same scenario, running
// of ideas here.

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// We want to be able to create factories dependent upon the
// settings that we want employees to subsequently be manufactured.
// Because, what we did in the previous example was that we basically
// created a NewEmployee to just specify some sort of flag, and then we customized it.

// It would be nice if we could do it in one statement.
// We want to be able to create factories for specific roles
// withing the company and there are two ways of doing it.

// There's the functional approach and the structural approach.

// -> Functional

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// <- Notice we're not creating an employee we're creating an
//    employee factory that we can subsequently use to fine tune
//	  those details of that object.

// One advantage is that now we have factories stored in variables and
// we can also pass these variables into other functions. Well, that's
// the core of functional programming. Functions into Functions Lj !

// That other approach is that we make that facotry into a struct.
// It's not strictly speaking necessary, at least not in Go, but we can do it.

// So if we want to somehow incorporate information about the fact that
// a particular factory initializes an employee with a particular position and
// annual income we would do something like this.

// -> Structural

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (e *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, e.Position, e.AnnualIncome}
}

func NewEmployeeFactoryStruct(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

// The only real advantage here is that after these factories are created
// the functional factories can't really be customized afterwards.
// But we can do this with structural factories, they actualy store fields.

// In terms of usability in third party code, let's say providing
// a function and passing it into some other piece of API, then the
// situation is different.
// Because obviously in these cases those are just ordinary functions and
// and passing functions into something is easier than passing in a specialized object.
// For specialized objects, whoever is consuming that object has to explicitly know that
// there is some sort of create method and then they have to call this create method.

// This is a situation for example where we might try to introduce some sort
// of interface which tells us explicitly that there's a create method and here
// are the arguments and then we could also use this to pass an interface of the
// factory, rather then the factory itself.

// Which ever option we go, it's just fine.
// I just had to much free time here.
// But the first one is probably more idiomatic, and would go for that one.
// And the second one is here if we need it for some reason.

func main() {
	// NewEmployee(1)
	// e.Name
	developerFactory := NewEmployeeFactory("dev", 175)
	managerFactory := NewEmployeeFactory("good for nothing", 175000)

	dev := developerFactory("Vincent")
	mng := managerFactory("Ho Chi Minh")

	fmt.Println(dev)
	fmt.Println(mng)

	bossFactory := NewEmployeeFactoryStruct("CEO", 1000000)
	bossFactory.AnnualIncome = 9000000000
	bss := bossFactory.Create("Bob")
	fmt.Println(bss)
}
