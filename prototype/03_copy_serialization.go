// Copy Through Serialization

// Previously, we've looked at this idea of defining
// a Deep Copying methond on every single struct that we
// have in our program or every single struct that we need
// to replicate.
// And then invoking that in a kind of recursive fashion.

// We also talked about the fact that it doesn't save us from
// handling all those structures which we don't have a DeepCopy method
// like slices for example.
// And the fact that we would have to perform our own copying operations
// relevant to that particular data structure.

// So we might have a question, well is there some magic that we can
// just *pooft* and have all of it go away and have a good replicaiton
// of objects where we just say let's copy this struct and that struct
// is copied including all dependencies including the following of the
// pointers and the rest of it.

// And luckily, the answer is yes.
// This is done using -> Serialization

// For now, we're going to use binary serialization,
// although of course, we could also serialize to something
// else like JSON.

// The idea here is that serialization constructs are very smart.

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// If we give them something like a Person, the serializer there is
// going to figure out that we have a string which has to be saved as
// a string.
// But also that we have a pointer there and we need to follow
// that pointer to take a look at the data stored at that location and serialize
// that data and not just the pointer value.

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// So serializers know how to unwrap a structure and serialize all of its members.
// And if we think about it, if we serialize the Person, to let's say a file or
// just to memory, we save all of it's state, including all the dependencies.
// And when we deserialize it, we construct a brand new object initialized with
// those same values.

// So effectively, we can have a Deep Copy performed automatically and so we can
// get rid of all those methods we had.
// So let's have a utility now, a DeepCopy method on Person.

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)

	return &result
}

// So, it looks like we've finally solved this tedious problem of
// copying objects and then of course there is the small matter of
// discussing the prototype design pattern.

// Because the prototype design pattern is all about taking a pre-configured
// object, like John there, making a copy and then customizing it like we do.

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
