// The Adapter Pattern

// Let's imagine we're working with some sort of API
// for rendering of graphical objects and let's  suppose that
// that API is completely vector based.
// All of the images are defined as a bunch of lines.

package main

import (
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

// Now let's imagine that a vector image is quite simply composed
// out of several of these lines.

type VectorImage struct {
	Lines []Line
}

// Suppose that we're consuming this API so we're being given this API
// by some external developer od system, and we have a fuction for making
// new graphical objects.

func NewRectangle(width, height int) *VectorImage {
	width = width - 1
	height = height - 1
	// <- Reason for this is things are tipically zero based
	// If we say we want an image that's of width 5 it has to
	// go from position 0 to position 4

	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

// ↑↑↑ We'll have to pretend that this is the actual interface that we're given.

// Now let's suppose that we can't work with this interface.
// And the reason could be is that we don't have a way of putting
// things graphically.

// Take our console for example, it can only show us rectangles by using characters.
// It can't really draw a proper rectangle.

// So let's suppose that the interface we have is somewhat different.
// ↓↓↓ The interface we have deals strictly in terms of points not in terms of lines.

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX++
	maxY++

	// preallocate
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// So now we can take a raster image and we can convert it into
// a string representation, which we can actually print out to the console.

// This introduces an obvious problem in the entire system,
// because we're given the interface where the only way for making a Rectangle
// is by making a Vector image, but unfortunately, the only way to print something
// is by providing a Raster image.

// @_@

// We need an adapter, something that takes a vector image and somehow adapts it
// into something which has a bunch of points in it so that those points can be fed
// into the raster image and sunsequently fed into the DrawPoints function.

type vectorToRasterAdapter struct {
	points []Point
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

// Now the issue here is how this thing gets constructed, because we're keeping
// it private, so we need some sort of factory function to actually construct
// a new VectorToRaster adapter.

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return &adapter
}

// Now we need a to take a line and decomposed it and set up a bunch of points.
// Before that we need minimax function.

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// ↑↑↑ This is a sad function. No ternary operator!

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := bottom - top

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("we have", len(a.points), "points")
}

// So this is how we build and adapter basically, so we've just
// built something which takes one API [VectorImages], and we've
// adapted it to a completely different API [RasterImages] that only
// wants to deal with points.

func main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	fmt.Print(DrawPoints(a))
}
