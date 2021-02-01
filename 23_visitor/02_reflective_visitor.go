// Visitor - Reflective Visitor

// In the previous example, we looked at
// an Intrusive Visitor, a visitor that causes us
// to modify existing structure.

// Now imagine if we wanted to do it differently.
// To concenrate all the print methods in a separate
// component, in a separate struct, or indeed just to
// a separate function.
// Imagine we didn't want to modify these types and give
// them additional methods.

// Can we do this?
// Well, it turns out we can.

package main

import (
	"fmt"
	"strings"
)

// First of all, we're going to Expression interface we had
// and delete the Print method.

// type Expression interface {
// Print(sb *strings.Builder)
// }

// <- This is what we had.

type Expression interface{}

// <- This is what we us now.

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

// Now, we're not going to have any separate struct
// for printing different types of expressions, instead
// we'll just have a function.

// And the kind of visitor we're going to build here is going to be
// what's called a -> Reflective visitor.
// What? Why is reflective?

// Well, because typically there is this construct called Reflection.
// That's when we look into a type and we look at what the type actually is,
// what kind of members it has, and Go has a certain amount of support for reflections.

// One of the trademarks of reflection is that we check the actual type.
// In the case of our expresion, but we don't really know what kind of
// an expression it is.
// Luckily for us, there are casts or we can try to do a cast in order to
// figure out whether this expression is a DoubleExpression, an AdditionExpression
// or something completelly different.

func Print(e Expression, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		sb.WriteRune('(')
		Print(ae.left, sb)
		sb.WriteRune('+')
		Print(ae.right, sb)
		sb.WriteRune(')')
	}
}

// <- So the idea is we have some expression and we print it
//	  inro a string builder.
// <- Notice that the handling of the printing of the different
//	  types is now handled in a separate function as opposed to inside
//	  each of those expressions

// Ok, now we have a self-contained function specifically for printing
// of the expression given a particular string builder, so we can start
// using that.
// And it works, same result as before.

// So this approach is somewhat better than the one we've
// taken in the previous demo, and the reason for that is
// because we've taken out this particular concern, the printing concern.

// We're following this idead of separation of concerns.
// We have a concern called printing, and we've taken it out
// and we've put it into a separate function
// We could have created a separate struct to have some sort of
// printer struct for keeping this function.

// Recap:
// -> The function we've used is a basically an implementation of
// 	  Reflective Visitor, because we're checking the type
// -> This approach is much better but it's not without it's downsize
// -> The most obvious one is what will happen if there is a third type,
//	  let's say substraction
// -> The reason why it's problematic is because all of a sudden we have
//	  to add additional code to Print function, and that would once again
//	  break the open closed principle
// -> We want to be able to extend things, we don't to modify existing code
//	  that's already been written and tested
// -> Another possible problem is what happens if we forget?
// -> If we forget to write that if else for a third expression it doesn't
// 	  get processed

// In the next example, we're going to take a look at the classic implementation
// it makes it pretty much impossible.
// It makes it impossible to set up a scenario where we suddenly forget
// to implement a particular support for a particular type of an expression.

func main() {
	// 1 + (2 + 3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}
