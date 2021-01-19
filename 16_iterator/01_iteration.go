// Iterator - Iteration

// Here, we'll talk about the general idea of iteration.
// To iterate something means to go through every single element,
// or maybe most elements, or selection of elements.

package main

import "fmt"

// We'll use a very simple example.
// Let's suppose we have a type Person.

type Person struct {
	FirstName, MiddleName, LastName string
}

// The question is: what if somebody wants to iterate all the names?

// One of the approaches is that we can simply expose an array.
// So we can take every single name, put it in array
// and just return it.

// And because arrays are iterable there's really no problem.

// This is obviously one way of doing it.
// But, it's not free, because we're effectively
// copying over a string so it's not as efficient.
// We might want to switch to pointers for example and
// there are also additional complications .

// Like let's suppose for example that if the middle name
// is empty we don't want to return it as part of the names.
// And ofcourse, it can be empty because some people don't have it.
// So how would we solve this?

// Here, in this particular paradigm it's difficult.
// We would have to jump from using an array to using a slice,
// but it's still manageable, it's still doable.
// A pain, in the rear, but still a possible approach.

// Now, another approach is to use a -> Generator.
// So this whole business of channels and Go routines,
// and all that stuff.

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()

	return out
}

// <- This was the second variety of iteration.

// And of course, for where two gathers a third one always appears.
// The third one is the most complicated variety of iteration, and
// that is when we use a seperate struct.
// This approach is very un-idiomatic [no hate Lj].
// It's the kind they use in C++, but we can also use it here.
// But we wont use it if realy realy don't need it.

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person: person, current: -1}
}

// ↑↑↑ Factory, just to initialize this person
// 	   -1 is the idea is that as we start iterating, we move
//		the current value

// And that's why we need a function to do that,
// and to return us a bool, indicating to us whether
// there is actually something to consume.
// Remember, we have an infinite number of names,
// we only have 3 of those.

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("We should not talk about this!")
}

// <- Ok, now that we have this setup,
// what we can od is we can use this iterator instead.

// Recap:
// ->	Typically when we talk about the iterator design pattern
//		we mainly talk about explicitly constructed iterators like -> PersonNameIterator
// -> 	We talk about seperate structures, which are used to track the position
// 	  	of where we're in the object that's being iterated
// 	  	and obviously we've appointed to that object so that
// 	  	we can go into it and get some information that we need
// ->	For a full implementation of an iterator, we'll take a look in
//		a next one

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func main() {
	p := Person{"Jean", "Henri Gaston", "Giraud"}
	// for _, name := range p.Names() {
	// 	fmt.Println(name)
	// }

	// for n := range p.NamesGenerator() {
	// 	fmt.Println(n)
	// }

	for it := NewPersonNameIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
