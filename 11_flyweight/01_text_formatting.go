// Flyweight - Text Formatting Example

// Here, we're going to take a look at a very simple example.
// Let's imagine we have some sort of text editing app.
// We're working with plain text but we want to add a bit of
// formatting, maybe we wanto to make text bold or italic or whatever.

// So we're goint to take a look at is we're going to take a look
// at a very inefficient version of doing that, and then we'll introduce
// a more efficient version which uses the Flyweight design pattern.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// So let's say that at the moment, the only kind of formatting
// we're going to have is capitalization.

type FormattedText struct {
	plainText  string
	capitalize []bool
}

// <- This boolean slice is very naive approach.

// Lets create a constructor first.

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText, make([]bool, len(plainText))}
}

// Now the real issue is how do we render this
// with all capitalization rules being applied?

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

// And we need some utility method to capitalize those letters.

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

// So lets have some bunch of junk text here.
// And it seems we have the right operation, that
// BRAVE word is being capitalized.

// But this approach is extremly inefficient!
// It's inefficient because we're specifying a huge
// boolean array with one element for every single character
// inside plaintext.

// The problem with this is that, if we imagine we're reading
// a text like "War and Peace", and we only want to capitalize
// a single word outside of thousands upon thousands of words we're
// going to be allocatin lots and lots of values that we don't even need.

// What we can do to remedy this is that we can introduce an
// idead of a text range which is simply the starting and ending
// positions of a range inside this rather large set of letters that we have.

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

// We can also have a utility method for figuring
// out whether the text range covers a particular point.

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

// So this way we're checking that the range is actually
// covering a particular position and we can use this subsequently
// when we make a different implementation of our formatted text struct.

type BetterFormatedText struct {
	plainText  string
	formatting []*TextRange // <- pointer here so we could manipulate those
}

// Let's have a constructor for this one.

func NewBetterFormatedText(plaintext string) *BetterFormatedText {
	return &BetterFormatedText{plainText: plaintext}
}

// Now, what we want here is that we want to be able to
// construct and return a range inside this text.

func (b *BetterFormatedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)

	return r
}

// <- We return the range here, so user can operate upon it.
// Take that range and customize it our heart's content.
// It can even be stored somewhere and used later. I don't care.

// OK, so the last thing with this setup that we need to do is
// to implement the stringer interface on BetterFormatedText so
// that we actually have a string method to work on.

func (b *BetterFormatedText) String() string {
	sb := strings.Builder{}

	for i := 0; i < len(b.plainText); i++ {
		c := b.plainText[i]
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}

	return sb.String()
}

// So what have we done here?

// Recap:
// -> We have essentially constructed a Flyweight -> TextRange
// -> It's an object that allows manipulation at a certain scale
//    but it's very compact, it tries to save memory

// Now the critical distinction between those two approaches
// is of course the savings in memory and that is what the Flyweight Pattern
// is actually for.
// It's a trick to try to avoid to much memory and to instead maybe
// introduce additional sort of temporary objects like we're doing with
// text ranges, but these temporary objects allow us to save a lot of memory
// which is always a good thing, right?

func main() {
	text := "This is a brave new world"
	ft := NewFormattedText(text)
	ft.Capitalize(10, 15)

	fmt.Println(ft.String())

	bft := NewBetterFormatedText(text)
	bft.Range(16, 19).Capitalize = true

	// ↑↑↑ When we call Range() a range is constructed and added
	//     to the set of ranger inside the formatting field of the
	//     better formatted text, but it also gets returned the client

	fmt.Println(bft.String())
}
