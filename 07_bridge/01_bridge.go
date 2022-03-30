// The Bridge

// Imagine that we are working on some sort of graphical application.
// And this application should be capable of printing different objects
// like circles, and rectangles and squares and...

// However it needs to be able to render them in different ways.
// And we, again, want to render those shapes in Vector form or in
// Raster form.

// If we were to let this situation out of control, we would have
// quickly end up with what's called -> Cartesian product complexity explosion.

// So if we have shapes like Circle and Square, and then we decide to render
// those in Vector and Raster forms, and if we plan this badly we'll end up
// with lots of different types.
// We would be stuck with:
// -> RasterCircle, VectorCircle
// -> RasterSquare, VectorSquare

// And if the number of types of those groups is larger than 2, then
// we would just have an explosion of all the different types.

// This must be simplified, to actually reduce the number of types
// that we need to introduce.

package main

import "fmt"

// Instead of having shapes specialized for different renderers
// we would take those and maybe introduce an interface like:

type Renderer interface {
	RenderCircle(radius float32)
}

// <- and we could have RenderSquare or RenderTriangle ...

// Now we can define types which take care of the rendering
// in a different format.

type VectorRenderer struct {
	// -> any kind of utility information related
	//    how we want vectors to be constructed
}

// And then we would have a function for actually rendering our type:

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a Circle of radius: ", radius)
}

// But similarly, we can have a Raster renderer which also knows how
// to render a cricle but it knows how to render it in a different way.

type RasterRenderer struct {
	Dpi int
	// ...
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for Circle of radius: ", radius)
}

// Now we can define the circle.
// It will be an ordinary circle, we won't define it like a vector
// or raster representation.
// We'll just have a circle which subsequently refers to or it has a
// -> Bridge to the renderer

type Circle struct {
	renderer Renderer
	radius   float32
}

// This Circle needs some functionality.
// Firs of all, let's have some sort of factory function.

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

// And then we need a function for drawing a circle:

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

// ↑↑↑ This is precisely where we would do these sort of specific
//     kinds of rendering

// But, we've defined a rendering member to actually render
// the things that we want.

// And also we could have other members like:

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

// Now, what we need to do to get everything to operate is
// we need to make a renderer which is a seperate component and
// then we provide that renderer into the circle.
// We introduce it as a dependency.

// In the end we've avoided this complexity explosion,
// That 2x2 setup of types that we might want to make,
// although not really.

// What this implies, and this is not the best of things,
// is that when we introduce a new shape, let's say triangle,
// we would by necessity introduce a new rendering method similar
// to the one that we have here.
// So the renderer interface would have RenderTriangle, and all the rest.

// And then, of course, because we want this interface to be implemented by
// both Vector and Raster renderer it would result in a cascading set of functions
// or methods rather being created on those renderers.

// But this is the price to pay for the additional flexibility and hopefully
// we can see that this is somewhat better then allowing a complexity explosion
// of an infinite number of types.

func main() {
	// raster := RasterRenderer{}
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}
