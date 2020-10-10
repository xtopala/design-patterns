// The Flyweight - Storage of User Names

// We're now going to take a look at the most
// classic implementation of flyweight design pattern
// and that's the storage of people's names.

// So why is it such a big deal?

// Let's suppose we have some MMOG which has a bunch of users.

package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

// We can have a constructor, which actually is a sort of
// factory function which initializes the User, and then we
// can start making these users.

func NewUser(fullname string) *User {
	return &User{FullName: fullname}
}

// So what's the issue here?
// Well. there's no problem if we have just 3 users.
// But, in the real world, out there if we take let's say 100 000
// different players in a online game or any online system for that
// matter we're going to have people with similar names.

// And every time we store one of those similar namses we're
// effectively duplicating memory, duplicating things.

// So what we could do here is we'll make a variable.

var allNames []string

type FrugalUser struct {
	names []uint8
	// ↑↑↑ making the assumption that there's only gonna be 256 unique names
}

// ↑↑↑ Now this new type of user is more frugal when it comes to
// the use of memory, because instead of storing the actual strings
// the names are going to be stored as unsigned integers.

// So this means that constructor for this new user type
// will be more complicated.

func NewFrugalUser(fullName string) *FrugalUser {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := FrugalUser{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

// Now of course, we have a problem here.
// If we wanted to get the full name of FrugalUser there is
// no full name as a single element, instead we have a bunch of
// integers inside names member.

// So whenever somebody needs a full name we have to basically
// provide a function which reconstitutes those.

func (fu *FrugalUser) FullName() string {
	var parts []string
	for _, id := range fu.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

// We get pretty much same output now as before and
// a bit more work has been carried out, but on the other hand
// there are certain memory savings, and the question is how much
// memory are we actually saving?

// So not bad, for couple of user we're saving a few bytes.
// In a really large scenario we would save huge amounts of memory.
// Because essentially, we're storing byte arrays of just two bytes per user.

// OK, now this is all fine, but what's the flyweight here?
// The flyweight is inside names member in FrugalUser.

// Just like in the previous example, where we had text formatting
// and we had that text range construct.
// These are also kind of like ranges, they're king of like pointers
// into all names here, except they're indicies into all names.

// So instead of operating on strings directly, we operate on
// integers which are representations or are pointers into this
// overall array.

func main() {
	john := NewUser("John Doe")
	amanda := NewUser("Amanda Hugandkiss")
	alsoAmanda := NewUser("Amanda Doe")

	fmt.Println("Memory taken by users: ",
		len([]byte(john.FullName))+
			len([]byte(amanda.FullName))+
			len([]byte(alsoAmanda.FullName)))

	frugalJohn := NewFrugalUser("John Doe")
	fmt.Println(frugalJohn.FullName())

	frugalAmanda := NewFrugalUser("Amanda Hugandkiss")
	frugalAlsoAmanda := NewFrugalUser("Amanda Doe")

	totalMem := 0
	for _, a := range allNames {
		totalMem += len([]byte(a))
	}
	totalMem += len(frugalJohn.names)
	totalMem += len(frugalAmanda.names)
	totalMem += len(frugalAlsoAmanda.names)

	fmt.Println("Memory taken by frugal users: ", totalMem)
}
