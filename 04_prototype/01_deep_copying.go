// The Deep Copying [Cloning]

// In order to be able to implement the prototype design factory
// we basically have to do one thing, we have to be able to perform
// called -> a Deep Copy of an object.

// So what's this all about, and why do we care?
// Let's just build a very simple scenarion, and see where it takes us.

package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

// So with this what we might want to do is we might want to make
// a Person and let's say that another Person lives buy, same building,
// we might want to perform a copy.

// But unfortunately this copying process isn't as obvious as one might think it is.

// Ok. John gets copied into Jane. OK, chill.
// So what do we expect here? Well the name gets copied.
// And also the address pointer gets copied, and this is a problem, because
// what we might want to do is customize Jane because that's the whole point after all.
// We want to give Jane a name.

// Changing a name should work, it's ok.
// But then we want to customize the address as well.

// We hit our obvious problem. They both have the same address.
// Address is ofcourse a pointer, and when we copied John to the variable Jane
// using a assignment, we also copied a pointer as opposed to made a new object
// a new address and copied the content.

// So that is the problem of -> Deep Copying
// And ofcourse, we have to hanlde this. But how?
// How to reliably copy John into a new Jane?

// Recap:
// -> Deep Copying is basically this idea that when we're copying the object
//    we're making copies of everything it refers to including all the pointers,
//    the slices and all the rest of it.
// -> If we don't do this then any object which operates as if it were a pointer,
//    like a slice for example, would be shared between the original object and the copy
//    and so modifying either of those would effect the other one as well. No bueno!

// However, the problem with this approach is that it really doesn't scale if we have
// a person which is composed of a Address and Address itself is composed of something else
// then we end up having this very complicated recursive structure and then we have to write
// lots of code just to be able to copy objects.

// So we probably want to organize this somehow and not do all of this work ;)

func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}}

	jane := john
	// jane.Name = "Jane" // this is fine
	// jane.Address.StreetAddress = "321 Baker St"

	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		jane.Address.Country,
	}
	jane.Address.StreetAddress = "321 Baker St"

	// <- now this is how we perform something called Deep Copying

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
