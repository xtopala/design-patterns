// The Adapter Caching

// One thing we need to be aware of is the creation
// of too many temporary objects.

// So for our previous case, in order to actually draw
// our lines as pixels, we turned every single line into
// essentially a bunch of points.

// That was alright, but it becomes a bit of the problem
// if we try to do it more than once.

package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

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

type vectorToRasterAdapter struct {
	points []Point
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return &adapter
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

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

// We can notice that when we run our code again, that there
// are operations that are not actually necessary, maybe we can
// live without those.

// The first set of operations on a 6 by 4 rectangle makes sense,
// we need those points but then we went ahead and we regenerated
// those points once again, because we made another adapter.

// We can avoid this.
// If we assume that our adapter is immutable, then it makes
// perfect sense to implement some sort of caching so we don't
// get this ridiculous duplication.

// The simplest thing we can do is build a very simple cache.

var pointCache = map[[16]byte][]Point{}

// Now we need to change our addLine() so that it doesn't add
// those points if they've already been generated.

func (a *vectorToRasterAdapter) addLineCache(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		for _, pt := range pts {
			a.points = append(a.points, pt)
		}
		return
	}

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

	pointCache[h] = a.points
	fmt.Println("we have", len(a.points), "points")
}

// Now use this in our adapter.

func VectorToRasterCached(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLineCache(line)
	}

	return &adapter
}

// And we can improve the situation even further by
// not storing those extra points by using Point pointers instead.

// Takeaway:
// -> If we generate temporary data, it makes sense for us to investigate
//    how we can make sure that this data isn't getting generated redundantly

func main() {
	rc := NewRectangle(6, 4)
	// a := VectorToRaster(rc)
	// _ = VectorToRaster(rc)
	a := VectorToRasterCached(rc)
	_ = VectorToRasterCached(rc)
	fmt.Print(DrawPoints(a))
}
