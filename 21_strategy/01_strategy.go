// Strategy

// So once again, we're going to do a demo where
// we're going to print a list of textural items and
// we're going to do this printing usign different formats.

package main

import (
	"fmt"
	"strings"
)

// We're going to introduce a type for our output formats,
// because we're going to have some sort of text processor which
// takes a list and it can print it usign either markdown or HTML.

type OutputFormat int

const (
	Markdown OutputFormat = iota
	HTML
)

// The idea is that we have some strategy for how to print a list.
// We can define a type for this.

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

// Now that we have this, we can build different strategies
// for constructing lists using the markdown format and using
// html format, and they're going to be significantly different.

// As a first, let's implement the markdown strategy.
// Which is gonna be an empty struct.

type MarkdownListStrategy struct{}

// <- Typically we would have some service information here
//	  or some formatting setting specifically for markdown.

// But all we're going to do here is implement the methods of
// the List Strategy interface.

// Now first thingto note is that when we're making lists of
// markdown elements, they're typically done in a way that there's
// no preamble, there's no start of the list and there's no end of
// the list like we have in html.

// Because of that, both Start and End methods are going to be empty.

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {}
func (m *MarkdownListStrategy) End(builder *strings.Builder)   {}

// Everything that has to happen happens when we're adding items.

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

// That's all there is to it, and now we can implement
// the HTML list strategy.

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("	<li>" + item + "</li>\n")
}

// This is how we can both write markdown as well as html.
// And now let's imagine that we have some sort of text processor.

// The text processor is this configurable component which we can
// feed this component to a list and specify the target strategy that
// we want to take.

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

// Now, let's make a constructor for this.

func NewTextProcessor(ls ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, ls}
}

// Notice that when we got started we defined a bunch of constants,
// under OutputFormat, and we can use these to force a switch from one
// strategy to another, so that means that somewhere inside the text processor
// we would have a method.

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}
	case HTML:
		t.listStrategy = &HtmlListStrategy{}
	}
}

// Now let's have a methond on the text processor where we take
// a bunch of items and we append them using the selected strategy.

func (t *TextProcessor) AppendList(items []string) {
	s := t.listStrategy
	s.Start(&t.builder)
	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}
	s.End(&t.builder)
}

// Let's also add a Reset method, because we want to be able
// to reset the internal strings builder.

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

// And let's have a string representation where once again
// we'll just going to implement the stringer interface.

func (t *TextProcessor) String() string {
	return t.builder.String()
}

// Finaly, we can take a look how all of this works.

// Recap:
// -> This was an illustration of how strategy works
// -> Essentially what we do is we have a member which can
//	  be defined to different kinds of structs
// -> In this case we have a List Strategy which is just an
//    interface, and can be implemented by either the markdown
//	  or html list strategy
// -> We can switch from another, and we can have it defined at
//	  initialization

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	tp.Reset()
	tp.SetOutputFormat(HTML)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)
}
