// The Dependency Inversion Principle

// This principle doesn't have anything to do directly
// with dependency injection.
// Those things need to be kept seperate!

// DIP states 2 things.
// -> High level modules should not depend on low modules
// -> Both of those should depend on abstractions

package main

import "fmt"

// Imagine we're doing some geneology research and we want
// to model relationships between different people.

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

// And then we can model different people

type Person struct {
	name string
	// DOB
	// ...
}

// To model a relationships between different people we could use some Info

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// And now we want to have all the data about the relationships between different people

// Low level module
type Relationships struct {
	relations []Info
}

// <- From this we could find all the children of particular person.

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		parent,
		Parent,
		child,
	}, Info{
		child,
		Child,
		parent,
	})
}

// And then we want to perform some research on this data.
// This Research, in our case here, is what we would call -> High Level Module,
// and Relationships <- Low Level Module

// Low level is usually some storage or persistance mechanism.
// And high levels are design to operate on some data.

// High level module
type Research struct {
	// break DIP
	relationships Relationships
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called: ", rel.to.name)
		}
	}
}

// <- Major problem here!

// ATM, the Research module is actually using the internals of the Relationships module.
// Relationships is a low lvl mod and it's using literally it's slice to get data.
// Now imagine if Relationships decides to change the storage mechanics,
// from a slice to lets say some database.

// -> All of the code which depends on that low lvl mod actually breaks.
// -> DIP is trying to protect us from these situations where everything breaks down.

// In actual fact it could be argued that the finding of a child of a particular person
// is something that needs to be handled not in higher-up, but in low lvl mods.
// Essentially, if we know the storage mechanic we can do an optimized search [like in a DB].

// Let's make this better

type NewResearch struct {
	browser RelationshipBrowser
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

// Now all the finding of the children is actually put into the low lvl mod
// and then we can rewrite the high lvl mod to depend on an abstract, just like we wanted.
// We depend on the Relationship Browser and then the investigation becomes different.
// Because the search for the children is now in low lvl part, and all we have to do
// is we have to handle it somehow.
// We need to perform the actual investigation.

func (r *NewResearch) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called: ", p.name)
	}
}

// We can notice that this became much more simpler.
// And the low lvl implementation details are not exposed to the high lvl research mod.
// Browser abstract finds all the children for a particular parent.

// Recap:
// -> High level modules should not depend on low level modules
// -> Both of them should depend on abstractions [Interfaces]
// -> We are protecting ourselves against changes

// If we decide to change the storage mechanic of relations
// from a slice to something more sophisticated then we would
// only be modifying the methods of Relationships.
// We would not be modifying the methods of for example Research
// because it doesn't depend on the low level details.

func main() {

	parent := Person{"John"}
	child1 := Person{"Chriss"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{relationships}
	r.Investigate()

	nr := NewResearch{&relationships}
	nr.Investigate()
}
