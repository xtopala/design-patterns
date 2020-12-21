// Proxy - Virtual Proxy

// A Virtual Proxy is a kind of a proxy
// that pretends it's really there when it's not necessarily.

// Let's imagine that we have an interface called Image,
// because image is something that we can draw, and in our case
// we're just going to emulate this process of drawing and image.

package main

import "fmt"

type Image interface {
	Draw()
}

// So now we might want to have a bitmap.
// We load the bitmap up from a file and then we draw it,
// so we can have a type for that.

type Bitmap struct {
	filename string
}

// Ok, so we can construct a bitmap and we can draw it.

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename) // just pretending, not really drawing
}

// Let's have a constructor for a bitmap, which simply
// initializes it with a file name.
// But in addition, what we're going to do is we're going to specify
// here that this is the point where we actually load the image from some file.

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading the image from", filename)
	return &Bitmap{filename: filename}
}

// We obviously need the image to actually construct the bitmap.

// Now let's imagine that somewhere down below we have a
// function for actually drawing some images.
// We pass it the image interface and we get to draw that image.

// Not really gonna draw it, just output a bunch of diagnostic calls.

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

// So far so good.
// What we can do is we can make a bitmap and
// we can feed that bitmap into this draw image function.

// This scenario has a problem.
// What happens if we never draw the image in the first place, that is
// if never actually invoke DrawImage function.

// So we have the invocation of new bitmap, and that's it.
// Now if we run this, there is a fairly obvious problem that
// we are still loading the image even though we never draw it.

// One attempt to fix this might be to introduce some kind of
// lazy bitmap, the kind of bitmap where the image doesn't get loaded
// until we actually need to render it.

// We can implement this using a proxy.
// A lazy bitmap is something that is going to wrap and ordinary bitmap.
// And it's also going to implement the image interface and
// provide the draw method.
// But it's going to do it differently.

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

// <- It's also going to reuse the underlying bitmap functionality,
//	  because we don't want to reimplement it again.

// The idea is that when we make a constructor we don't specify
// the bitmap yet because this bitmap right here is going to be
// lazily constructed.

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

// <- It's only going to be constructed whenever somebody needs it.

// Now we can implement the image interface, and
// we obviously want to use the underlying bitmap, but
// we need to make sure that it's constructed because ATM the
// pointer has a value of nil.

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

// And this time the order of the output is going to be different.
// We don't load the image prematurely.

// We can also draw image twice and loading only happens once,
// since it's lazy.

// Recap:
// -> The demonstration here shows how we can build something
//	  typically called a virtual proxy
// -> The reason why it's virtual is because when we create a lazy
//	  bitmap using the new lazy bitmap function it hasn't been materialized yet,
//	  meaning that the underlying implementation of the bitmap hasn't even been
//	  constructed and it's only constructed whenever somebody explicitly asks for it

func main() {
	_ = NewBitmap("demo.png")
	// DrawImage(bmp)

	nbmp := NewLazyBitmap("lazy-demo.png")
	DrawImage(nbmp)
	DrawImage(nbmp)
}
