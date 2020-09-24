// The Composite Pattern - Illustration through Geometric Shapes

// If we go back to that typical drawing application,
// where we can draw all sorts of different vector shapes
// we know that we can take a shape and we can drag it aroung
// the screen, but we can also take several shapes and we can drag
// them together around the screen as a single group.

// Let's try to implement this scenario.

package main

import (
	"fmt"
	"strings"
)

// We'll need some object, that can be one of two things.
// Either it's a simple object [square, rectangle] or something
// that is a combination of different objects [children].

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

// Now we need to print this, and printing this to a console
// would be completely different depending on whether it's just
// a single object or whether it's a graphic object container with
// a bunch of children.
// Since this is a recursive relationship, we could have a drawing.

// In order to print this our, we'll define a String method:

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

// And an utility print method, which tracks the depth of our recursion,
// because we're going to go into objects and objects and objects ...

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

// Let's say we want to play with some squares and circles:

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

// With all of this we can now setup a scenario where we have
// a drawing which is a graphic object and we also have these
// circles and squares on it.

// Takeaway:
// -> We can have these objects like a graphic object of infinite depth,
//    because a graphic object can contain any number of graphic objects to
//    infinity and so a graphic object can be treated as a single shape or a scalar.

// So for an example, when we have a new Circle or new Square, we're
// effectively making a scalar, we're making a singular object not a
// collection of objects.

// -> If we decide to manipulate the children of an object, then
//    all of a sudden, it becomes a collection of objects.
// -> And that is the goal of the Composite Design Pattern, it
//    doesn't really matter whether or not our object is scalar object
//    or collective object which is composite.

// Ultimately we can write algorithms like that print() method
// which don't really care, because all they do is they perform
// the appropriate check.

func main() {
	drawing := GraphicObject{"My Doodle", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle("Red"))
	drawing.Children = append(drawing.Children, *NewSquare("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))

	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}
