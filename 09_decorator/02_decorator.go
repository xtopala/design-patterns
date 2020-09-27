// The Decorator

// We're going to use a bunch of geometric shapes
// for our example here, and we'll want to extend
// the functionality of those geomtric shapes by giving
// them additional properties.

// We need some scenario for this.

package main

import "fmt"

// So let's suppose that we have some sort of interface
// [Shape] and this interface is going to allow a shape
// to render itself.

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius: %.2f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side: %.2f", s.Side)
}

// Now imagine we these shapes operating in our system,
// and what we want to do is we want to color them.

// One approach is that we jump back into those shape structs
// and we add additional members.
// However, the real problem with that is it breaks the Open Closed Principle.

// So do we extend these types?
// We could simply aggregate and thereby extending the type.
// And if we had a small number of structs that we want to extend
// this could pretty well work.
// But if we had lots of different structs then having a counterpart
// colored type for every single one of those just doesn't make sense.
// It's simply too much effort, and without generics we can't specify that
// Shape type is going to include some provided type as type parameter.

// Instead we could use Shape in there and we can say that
// we'll have a ColoredShape and that's going to be our first -> Decorator.

type ColoredShape struct {
	Shape Shape
	Color string
}

// <- Now this one also needs to implement the Shape interface.

// And we can use underlying implementation of the interface
// so we can use it's Render() method

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color: %s", c.Shape.Render(), c.Color)
}

// <- This is now our firstborn Decorator, and we can start using it.

// Ok, this decorator might seem really wonderful but,
// there are certain things that we lose.

// We'll remember that on our Circle we've defined the Resize() method.
// When we're operating on Circle, that Resize is not a problem.
// But the problem is once we've made a decorator, once we've put
// ColoredShape over the ordinary shape, what we can't do is we can't
// use that Resize(), because it's no longer available.
// And there's no real solution to this, because we're not aggregating
// anything, we've lost that particular method.
// The only way we can restore it is to add it again.

// But here's the problem!

// The problem is that, how do we add it without also adding it
// to the interface, because remember, it's only the Circle type
// that has the resize method.

// That's a real life limitation of the Decorator approach.

// But one upside is that decorators can be composed.
// Which means we can apply decorators to decorators.
// No problem in doing this, no problem, at all.

// So lets make another one. A TransparentShape.

type TransparentShape struct {
	Shape        Shape
	Transparency float32 // <- this will be normalized into some int value
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

// <- We can use this one over the ColoredShape

// And there we go, decorators can be composed.
// But this does not do any kind of detection, in terms of
// circular dependencies or in terms of a repetition.

// For example, we can apply a ColoredShape to another ColoredShape
// and that's not going to be a problem.

// If we want to start detecting those things it's going to be
// a lot more work, and to be honest, not sure if it's really worth it.

func main() {
	circle := Circle{2}
	circle.Resize(2)
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
