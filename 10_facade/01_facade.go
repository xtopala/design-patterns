// The Facade

// The idea here is that we can hide a very complicated
// system behind a very simple interface.

// Unfortunately, we don't have time here to build
// a complicated system but we use a more real life
// example of something, and that is the idea of multi-buffer
// -> Multy-board terminal.

// So we would have different viewpoints which in turn
// are attached to different buffers, but we still want to
// work with a console and in as much as simplified manner as possible.
// We want a simple API over something that's complicated.

package main

import "fmt"

// So if we were to build such a thing, we would start with
// making a buffer.

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height, make([]rune, width*height)}
}

// <- So this is how we could initialize the Buffer
// and have a factory function which constructs it.

// And then we could have some utility method for
// getting a character at a particular possition in the buffer.

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

// So this is one of the components of our rather complicated
// system, and the of course, we need to present this Buffer on
// the screen.
// And remember, a buffer can be really large.
// It can be hundreds of lines long if we want to preserve
// the history of the inputs to a particular terminal, but
// we can only show a part of that buffer, and for that we have
// a new construct -> the Viewport.

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

// And in the followint fashion, we can have another
// utility method for getting a character at a particular position,
// incorporating the knowladge about this offset member of ours.

func (v *Viewport) GetCharacter(index int) rune {
	return v.buffer.At(v.offset + index)
}

// <- This way we get the character from the start of
// the visible area as opposed to the start of the entire buffer.

// So now we have a situation where we have Buffer and Viewport
// and we can imagine a Console, multi-buffer console, being a kind
// of combination.

// We would have a lots of viewports and lots of buffers,
// but we also want a simple API for just creating a console,
// which contains all of these, behind the scenes.
// And this is where we could build -> a Facade.

type Console struct {
	buffer    []*Buffer
	viewports []*Viewport
	offset    int
}

// And what we're going to have here is we're going
// to have an initializer which creates a default scenario.
// Now that is where a console has just a single buffer and a viewport.

// And if we look at terminals in Windows, Mac, and Linux
// the default implementations of a console is the one where
// we have just one buffer and one viewport.

// Not very exciting, but this is the kind of
// simplified API that we would expect a facada to provide.

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

// And once again, now that we have a Console, we can
// also have a high-level get character at kind of function,
// for figuring out a character's position at a particular point
// in the console.

// In order to do this we have to grab a buffer and look inside it.
// And for grabbing that buffer we might want the viewport,
// because remember a viewport has this get charcter at function.

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacter(index)
}

// This way we would use this instead of working with
// low-level constructs like buffers and viewports.

// Recap:
// -> The idea of Facade is basically providing a simple API
//    over something that's complicated
// -> Here, if we wanted to work with buffers and viewports manually
//    we would have to manage the ourselves, and that's a lot of work
// -> In most cases, we need an implementation where there is just a
//    single buffer and just a single viewport and that's exactly what
//    we are providing as part of part of the new Console factory function

// And of course just because we're making a Facade doesn't mean
// that we have to obscure the inner details.
// So if somebody wanted to mess about with the buffers in the viewport
// they can do it through the console.
// Or if they want they can use the viewport type and buffer type directly
// without even working with the console.

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
