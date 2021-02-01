// Visitor - Intrusive Visitor

// Aigth, we're going to take a look at different
// implementations of the visitor design pattern.
// And the first one that we're going to examine is
// what's typically called an -> Intrusive Visitor representation.

// So let's take a look at how this can actually work.

package main

import (
	"fmt"
	"strings"
)

// Now in this scenario we're going to be taking a look
// at exoressions such as (1 + 2) + 3
// We'll be having numeric expressions but they will be
// represented in terms of structs, and what we mean is that
// we'll have some sort on an expression interface.

// type Expression interface {}	<- not using it, we have new one

// <- And we can think of expressions as being composed of
//	  other expressions

type DoubleExpression struct {
	value float64
}

// <- We're calling it Double because we're using double precission

type AdditionExpression struct {
	left, right Expression
}

// With this setup, we can construct such an expression.
// Now that we have left and right parts, what we want to be
// able to do is we want to be able to print an expression for example.

// But how can we do it?
// Because, remember this is an abstract syntax tree.
// And we have to traverse this tree in order to be able
// to do anything, but where is the printing done exactly?

// So the first implementation of Visitor is an Intrusive one.
// And what we mean by that is whenever we call something intrusive,
// it basically means that it itrudes into the structure that we've
// already created.

// And any intrusive approach is by definition a violation of
// the Open Closed Principle, because remember that it basically
// states that once we have defined a Double Expression and an Addition
// Expression we shouldn't be able to jump back into those structs and give
// them new methods, for example.

// That's one of the things we're trying to avoid but since we're
// doing an intrusive approach we're going to modify this entire hierarchy.
// We're going to modify the interface.

// We tell the interface that it knows how to print itself.

type Expression interface {
	Print(sb *strings.Builder)
}

// And then we're going to have implementations of Double Expression
// Addition Expression.

func (d DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

func (a AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.left.Print(sb)
	sb.WriteRune('+')
	a.right.Print(sb)
	sb.WriteRune(')')
}

// This seems like something that'll work.
// And it did, cool.

// Recap:
// -> On the one hand we managed to implement a kind of visitor
// -> Now one question we might have is who exactly, or which component
// 	  exactly is the visitor?
// -> Which component is actually visiting every single element of this
// 	  abstract syntax tree?
// -> And what allocates the visitor here is the strings Builder,
// 	  so the string builder is the one that gets passed into every
// 	  single print method
// -> It gets to visit addition expressions as well as double ones,
// 	  and that's why it's called a visitor
// -> On the other hand, we've said that this visitor is intrusive,
// 	  so it's not the best visitor that we can build
// -> Implementing this one implies that we modify the behaviour of
//	  both any interfaces that we have as part of the element hierarchy as
//	  well as the elements themselves
// -> Every single element suddenly has to have this addditional method

// Aftermath:
// -> Imagine for a second that we want to have another visitor
// 	  that actually calculates the value
// -> Unfortunately in the setup, what we'd have to do is we'd have
//	  have to go into the Expression interface, and add another method,
//	  called Evaluate maybe, which returns float64 and then we would have
//	  to implement this in both DoubleExpression and AdditionExpression
// -> Another very important concept that we need to cover here is the idea
// 	  of separation of concerns and single responsibility
// -> It's possible to say that it's kind of the responsibility of each expression
//	  to print itself, but not necessarily
// -> It would make more sense if we had a separate component, a separate struct
//	  which knew how to print DoubleExpression, AdditionExpressions as well as any
//	  other kind of expressions we would have as part of our program

// <- This is what we're going to take a look at in the next example,
//	  and we'll try to take out those print methods from expressions and
//	  just have them exists in a separate component

func main() {
	// 1 + (2 + 3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	sb := strings.Builder{}
	e.Print(&sb)
	fmt.Println(sb.String())
}
