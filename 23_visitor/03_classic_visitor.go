// Visitor - Classic Visitor

// Last time, we've looked at the
// Reflective Visitor, where the idea was
// that we take some general expression and then
// we perform a bunch of typecasts and then attempt
// to figure out what kind of expression we actually got.

// So if we are going to implement the Classic Double Dispatch
// implementation of the visitor, we need to discuss what dispatch is.
// Because dispatch is a very strange word.

// We might have not heard of it, but basically the idea is that
// dispatch is all about figuring out which function or which method
// we're actually going to call.
// In some cases it's easy.
// But in other cases it's pretty much impossible.

// Let's look at a very simple example, shall we?

package main

import (
	"fmt"
	"strings"
)

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

// Let's suppose that instead of having our Print function,
// from before, we decide to have a separate Print function for
// each type of expression.

// For our Double Expressions we would have something like this.

// func Print(e DoubleExpression, sb *strings.Builder) {
// 	sb.WriteString(fmt.Sprintf("%g", e.value))
// }

// And for Addition Expression, something like this.

// func Print(e AdditionExpression, sb *strings.Builder) {
// 	sb.WriteRune('(')
// 	e.left.Print(sb)
// 	sb.WriteRune('+')
// 	e.right.Print(sb)
// 	sb.WriteRune(')')
// }

// We can try and od something like this.
// But of course, it's not going to work.
// We all probably know why it's not going to work,
// and there are actually multiple reasons why.

// One of the reasons of course is that we can't overload functions.
// We have function Print declared at two places, and if we
// try to compile this we'll get an obvious message.
// We're not allowed to do this.

// But just for a second, imagine that we're allowed to.
// And we overload our function, would the code work then?
// The answer is no, the code would still not work.

// The reason why is because of the problems we have
// at the Print function itself.
// Whenever the compiler encounters either e.left it
// knows that e.left is an expression.
// It doesn't know that e.left is an Additiona Expression or
// a Double Expression and it can't figure it out at compile time.

// And compiler wants to know the static type of e.left
// in order to be able to do something with it, to call a
// particular function, but because it doesn't that's pretty
// much impossible.

// If we were to call some function which has a way of
// accepting a Double Expression this wouldn't work because
// it's not a Double Expression it's just an ordinary expression.

// So this is the limitation of Dispatch.

// Since Dispatch is all about choosing which function to call and
// at the moment because e.left and e.right are expressions we can't
// make any choices.

// That's the exact reason why this idea of Double Dispatch is used.
// So instead, we take e.left and then we call something on it,
// some sort of an Accept method.

// Now inside this Accept method we do know who the owner is,
// because that's the receiver parameter.
// And now we can jump from an Accept method back into some sort
// of Visit method, like Print would be if it was a method, for example.
// By performing this double jump, we implement something
// called -> Double Dispatch.

// Double Dispatch is being able to choose the right method,
// not just on the basis of some arguments, but also on the basis
// of who the caller is.

// So, we're going to try to implement this whole thing.

// First, we need to modify the Expression interface.
// And somebody might say, well hold on hold on.
// We've just spent so much effort telling everybody about
// Open Closed Principle and how it implies we don't do this.

// Well the compromise here is that we can modify this interface
// but we only do it once.
// And then that operation is leveraged for many different
// kinds of visitors.
// And that operation is called Accept().

type Expression interface {
	Accept(ev ExpressionVisitor)
}

// <- ExpressionVisitor is nothing more than a interface

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

// Now that we have this interface, what we can do is
// if want to make let's say a printer for example we need
// to implement it.
// But before we do that we have to take the expression interface
// and implement specifically the accepth method in both Double
// Expression as well as Addition Expression.

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

// <- We return control back to the Visitor, but we return
//	  control specifying what kind of expression we actually have.

// So, we're invoking the right method and we do the same thing
// for the Addition Expression.

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

// And this is how we perform a double jump.
// And we need some demo, we need some sort of printer.

type ExpressionPrinter struct {
	sb strings.Builder
}

// When it comes to visiting the different kinds
// of expressions what we can do is we can implement
// that visitor interface.

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep) // <- this is the place where double jump magic happens
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}

// ↑↑↑ We need to try and explain this a little bit.

// So we call e.left.Accept() and the reason why we can
// call this regardless of what e.left is because either left
// is an expression and an expression is an interface that defines
// a method called Accept(), so we know that Accept is there.

// We also pass ExpressionPrinter to Accept, we pass it the visitor.
// So we go into except for the left side.
// And we either end up in the Accept for Double or Addition Expression.
// Depending on where we end up, we either call the VisitDoubleExpression
// or VisitAdditionExpression on the Visitor.

// When we call one of these methods we end up back where we started.
// But, we end up in the correct overload with the correct information
// about whether we have a Double Expression or indeed we have an Addition Expression.

// Now that we have all of this,
// let's just make a constructor for the new Expression Printers.

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

// Let's also implement the stringer interface on
// an Expression Printer.

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

// We can finally see how our previous example will change,
// on the basis of this Double Dispatch Visitor.

// Now, we need to talk about the sensibility of this
// approach because we've been fighting for having the support
// of the Open Closed Principle, so the idea is that things are
// open for extension but closed for modofications, but still some
// modifications might be required.

// Like imagine we also have a Substraction Expression,
// then unfortunately we would have to go into the ExpressionVisitor
// interface and we would have to implement a VisitSubstraction expression.
// But as soon as we did this, an interesting thing would happen.

// For all of the visitors which we wrote, it would be mandatory
// for them to support the Substraction Expression.
// And this is a drastic difference to the previous example
// where we could actually forger to handle a Substraction Expression
// and everything would still pass.

// Recap:
// -> In this particular scenario we can't forget to handle
// 	  a particular type of expression
// -> That is the modification that we would have to make in
//	  both the interface as well as the implementers of this interface
//	  like for example the expression printer
// -> On the other hand the situation when you need a new visitor
//	  is much better, it's easier and it does follow the OCP
// -> This approach is the most common kind of visitor we're likely
// 	  to see
// -> And the best benefit it provides is the flexibility as we create
//	  additional visitors

// Bonus:
// Imagine if we want to not only print the expression,
// but also give it a value.
// Evaluate it's final value.

// It's very easy for us to write an additional visitor.

type ExpressionEvaluator struct {
	result float64
}

// And once again, we need to implement the Expression Visitor.

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ee.result = e.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
	e.left.Accept(ee)
	x := ee.result
	e.right.Accept(ee)
	x += ee.result
	ee.result = x
}

func main() {
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)
	// fmt.Println(ep.String())

	ee := &ExpressionEvaluator{}
	e.Accept(ee)

	fmt.Printf("%s = %g", ep, ee.result)
}
