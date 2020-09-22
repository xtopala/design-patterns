// The Builder Facets

// In most situations, throughout our daily jobs,
// a single builder is sufficient for building up our peculiar objects.
// But nothing is that simple, and there are situations where we need more
// than one builder, where we need to somehow separate the process of building
// up the different aspects of a particular type.

// So this could work in this way.

package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string
	// job
	CompanyName, Position string
	AnualIncome           int
}

// So imagine we want to have a separate builders for building up
// the address information and for building up the job information.
// So the big question is how do we do it?

// We have to have a type that is a starting point.

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// Now what we could do is have additional builders for the address and job.
// These additional builders will aggregate the person builder.
// Now soon as we do this, we automatically get a pointer to person.

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// As our starting point was PersonBuilder, we want to be able to provide interfaces
// which are provided by address and job builders, respectively.
// We need some utility methods for this to happen.

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

// So now we have ways of transitioning from a Person builder to either
// a PersonAddressBuilder and a PersonJobBuilder. But we have to realize that,
// in effect a PersonJobBuilder and PersonAddressBuilder are both PersonBuilders. @_@

// As a result when we have a PersonAddressBuilder we can use the Works method to jump
// to a PersonJob builder, and vice versa.

// So what we could do now is populate those methods.

// Address Builder
func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

// Job Builder
func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnualIncome = annualIncome
	return b
}

// In a way, we've setup a kind of tiny DSL, a tiny domain specific language
// for building up a person's information.

// And now of course, we have to somehow provide a build method where we can
// actually yield the whole thing.
// So once the object is build up we return it, so let's add it.

func (b *PersonBuilder) Build() *Person {
	return b.person
}

// Recap:
// -> Instead of using 1 builder we have 3 :[
// -> PersonBuilder by itself doesn't do anything, appart for giving us 2 builders
// -> But we can switch from one to completely different one, because they both happens to be
//    the same build [PersonBuilder]
// -> This demonstrates that it's possible to have several builders who are working together,
// -> Only requirement is that they all agregate the same type which actually shares the pointer
//    of the object being built up

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Extra").
		AsA("Poor Dev").
		Earning(10)

	p := pb.Build()
	fmt.Println(p)
}
