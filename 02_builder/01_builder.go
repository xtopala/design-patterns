// The Builder Design Pattern

// We're going to begin here by using a builder that's
// actually already built into Go. And that is -> The String Builder

// So we have a situation where let's say we're writting a web service.
// And web server is supposed to serve HTML. Classic based stuff here.
// It also could serve other things, like JavaScript, but who cares for JS.

// The obvious thing here is that we need to build strings of HTML
// from ordinary text elements, for example we have a piece of text and
// we want to turn that text into a nice looking piece of paragraph.

package main

import (
	"fmt"
	"strings"
)

// Here we could already use something like strings builder.
// This is already a part of Go SDK and it helps us concatenate strings together.
// It writes several strings one after another into a buffer and then
// get the final concatenated result for us.

// Now this so far was no problem, at all.
// But lets suppose we have a list of words and we want to put them into a list.

// We got what we wanted.
// But this whole process of building up a HTML by using these tiny little
// nuggets of text and using the string builder is a bit too much work really.

// We could take those encrusted little nuggets, and put them into structures.
// And this is the reason we get builders in a first place.
// The idea is we have a some sort of object that we want to build up, in steps.
// And we want to make it convenient, not to showell a way through all those nuggents.

// We can represent the entire HTML construct as a tree, as it should be,
// and the just learn how to print it and give the user a nice builder component
// where they can add elements to this tree without them necessarily even being
// aware that the tree is actually there. Wooah. Tripy!

const (
	indentSize = 2
)

type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	return sb.String()
}

// This HTML Builder now only cares about the root element.
// So long as we have a root element we can get the actual representation.
// We'll also cash the root name separately because sometimes we need to reset the builder.
// Wipeout everything in it, and purge with fire.

type HTMLBuilder struct {
	rootName string
	root     HTMLElement
}

// Now we need a utility function to create a builder.

func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{rootName: rootName, root: HTMLElement{
		name:     rootName,
		text:     "",
		elements: []HTMLElement{},
	}}
}

// But now we also crave an utility methods for acually populating HTML elements.
// And also we need to have a string representation for the HTML Builder itself,
// which is just the representation of the root element.

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

func (b *HTMLBuilder) AddChild(name, text string) {
	e := HTMLElement{name, text, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
}

// Now the end user just need to care about the utility calls.
// They don't care about anything else, not a scratch.

// So one thing that is interesting here, which shows up quite a bit
// inside the builder pattarn is the use of -> Fluent Interfaces

// Fluent interface basically is an interface that allows us to chain calls together.
// Now chaining calls in Go isn't really that convenient because of hanging dots,
// but I can live with that.

func (b *HTMLBuilder) AddChildFluent(name, text string) *HTMLBuilder {
	e := HTMLElement{name, text, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)

	return b
}

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("<li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// <- This is deadening

	// -> The better way

	b := NewHTMLBuilder("ul")
	// b.AddChild("li", "hello")
	// b.AddChild("li", "world")

	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b.String())
}
