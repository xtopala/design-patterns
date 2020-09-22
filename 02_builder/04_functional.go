// The Functional Builder

// One more way of extending some builder that we already have
// is by using an functional programming approach. Functional is bliss!

// So lets suppose that we have a simple Person type, a peasent, and
// we want to have a system to build that peasent up. Extend it.
// Into the higher society.

package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

// So here, instead of just performing the modifications straight away
// what we do instead is we add this modification to the list of actions.

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})

	return b
}

// We also need a method for actually building up
// the person once all those actions have been compiled.

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}

	return &p
}

// So the benefit of this setup is that it's very easy to extend the builder
// with some additional build actions without messing about with making new builders
// which aggregate the current builder and so on and so forth.

// We could also extend this further.

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})

	return b
}

// Since we have fluent interface here, it allows us to effectively chain our calls.

// So this whole setup ilustrates that effectively what we can do is
// we can have a kind of delayed application of all of those modifications.
// So our builder instead of just doing the modifications in place it can keep
// a list of actions to perform upon the object that's being constructed.

func main() {
	b := PersonBuilder{}
	p := b.Called("Gleb").
		WorksAsA("pig feeder").
		Build()
	fmt.Println(*p)
}
