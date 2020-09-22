// The Copy Method

// So now that we know what Deep Copying really is,
// we might want to somehow organize our code so it becomes easier.

// One very simple approach is that we take every single struct
// that we have in our model in we give this struct a method called
// DeepCopy which explicitly performs a copy.

package main

import "fmt"

// Here would have an Address for an example.

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

// Now lets add something else to our Person, just to demonstrate that
// this is also necessary for other types.
// Lets suppose that we have a list of friends.

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// So now we have to perform a Deep Copy on that as well, so we could
// also define a DeepCopy method on Person.

func (p *Person) DeepCopy() *Person {
	q := *p // 1st we copy everything that can be copied by Value
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)

	return &q
}

// So we perform the deep copy where every single collection
// of elements, every single pointer actually gets unwrapped and
// and gets copied correctly.

// So we can customize Jane once again.

// Recap:
// -> We can organize our own objects to have to have some sort of
//    DeepCopy method available on them
// -> However this still leaves open the problem of what to do with types
//	  which we don't know

// For example we had that slice there, and we can't just go ahead and
// add additional behaviors to that slice.
// Essentially we're stuck with having to call some sort of copy method
// or something to that effect.

// Or even if we could add the behaviors to the slice,
// how would that change things?
// It would still force us to basically double and triple check every
// single one of our structs and make sure that every single one of the
// member types has a DeepCopy method.

// This one is a workable solution. It works ok.
// But it's not ideal, we still have to do a lot of work in order
// to get the copying done.

func main() {
	john := Person{
		"John",
		&Address{
			"123 London Road",
			"London",
			"UK",
		},
		[]string{"Chirs", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
