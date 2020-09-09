// The Liskov Substitution Principle

// First off, this principle insn't applicable to Go
// Because it primaraly deals with inheritance.
// LSP basically states if we have some API that takes a base class,
// and works correctly with that base class, it should also work with derived class.
// But, in Go we don't deal with base classes and derived classes [thank you Gopher gods]

// Since this concept doesn't exist here, as it should be elsewhere,
// but we're going to try out a variation of LSP that does apply to Go.

// Let's suppose that we're trying to deal with geometric shapes of rectangular nature.
// And we decide to have an interface that allows the specifications of certain operations.

package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// Now this is all fine, and let us create a function that can use this.

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Println("Expected an area of: ", expectedArea, ", but got: ", actualArea)
}

// Still everything looks fine, until we decide to break the LSP.
// Lets imagine that we decide to be smart, and we make a type called Square.
// In our infinite wisdom, we decide that it aggregates a Rectangle and
// that Square is going to enforce this idea that width is always equal to height. Always.

type Square struct {
	Rectangle
}

// Because we're going to enforce this, we need some sort of constructor.

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size

	return &sq
}

// In addition what we'll do, and this is the part that breaks LSP, is
// that we'll have methods for SetHeight and SetWidth

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

// <- This is insidious!
// <- Prime example of violation!

// Let's observe the UseIt function, LSP states the following:
// If we're expecting to extend something up the hierarchy,
// in the Sized argument, it should continue to work.

// -> We took a Rectangle and decided to extend it to a Square.
// Square should also continu to work. But it does not in it's current form.

// LSP -> If we continue to use generalizations, Interfaces in our case,
// then we shouldn't have inheritors or shouldn't have implementers of those
// generalizations break some of the assumptions which are set up at the higher level.

// And it's one of those situations where there is no right answer. No right solution.
// We can take different approaches, and we can say that squares don't exist here.
// Since every square is a rectangle, we don't take kindly to those types around here.

// Or, we could explicitly make illegal states unrepresentable.
// So our Square doesn't really have width and height, it has a size and that pretty much it.

// Recap:
// -> The behaviour of implementers of a particular type should not break
// the core fundamental behaviours that we rely on
// -> We should be able to continue using Sized objects, instead of trying to
// somehow figure out, for example by doing type checks, if we have a Square or Rectangle

type BetterSquare struct {
	size int // width, height
}

func (s *BetterSquare) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)

	bsq := BetterSquare{2}
	UseIt(bsq.Rectangle())
}
