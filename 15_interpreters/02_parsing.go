// Interpreters - Parsing

// Part two of making an Interpreter is basically
// taking these tokens we got from Lexing and Parsing them.

// -> Parsing: turning tokens into more sophisticated structures

// More sophisticated structures can be traversed in a recursion
// and we can actually evaluate the numeric values of expressions.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// In order for all of this to happen we
// need to introduce a bunch of structs and interfaces.

type Element interface {
	Value() int
}

// ↑↑↑ This thing is going to return the value of either
//	   a simple construct, like a number, or a complicated one
//	   like a binary expression.

// Why binary?
// All that we have in our model is we have pluses and minuses,
// and those both take two operands, so they're binary operations.

// Before we do that, let's make a Integer type.

type Integer struct {
	value int
}

// <- This is a primitive for every single integer token

func (i Integer) Value() int {
	return i.value
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

// Now the interesting thing is binary operation.
// There's addition and substraction.

type Operation int

const (
	Addition Operation = iota
	Substraction
)

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

// So now, we want to be able to evaluate the value of this binary operation.
// We want to implement the Element interface, and
// get the value of the operation.

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Substraction:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("Unsupported operation")
	}
}

// With this whole setup, we need to write
// a new function called Parse, which will turn this set
// of tokens into a top level binary operation.

// And we're going to assume that every expression is
// a binary operation, as we tuned ourselves for that already.

// The most complicated part of all
// here is the left and right parentheses ( )
// When we encounter the left parenteses -> ( <- we're going to
// find a location where we encounter the right parenteses -> ) <-
// and then we're going to take everything in between and we're going
// to feed it recursively into the Parse method.
// Thereby, we're getting an element and that element
// is going to be stored.
// Phew!

func Parse(tokens []Token) Element {
	res := BinaryOperation{}
	haveLHS := false
	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case Int:
			// we're assuming this will always succeed
			n, _ := strconv.Atoi(token.Text)
			integer := Integer{n}
			if !haveLHS {
				res.Left = &integer
				haveLHS = true
			} else {
				res.Right = &integer
			}
		case Plus:
			res.Type = Addition
		case Minus:
			res.Type = Substraction
		// ( )
		case Lparen:
			j := i // yeah, this is outside, we need it later
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}

			var subexp []Token
			for k := i + 1; k < j; k++ {
				subexp = append(subexp, tokens[k])
			}

			element := Parse(subexp)
			// we have the same thing as before, we need to decide
			// wheteher or not this have the left hand side or right
			if !haveLHS {
				res.Left = element
				haveLHS = true
			} else {
				res.Right = element
			}
			i = j
		}
	}

	return &res
}

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
	var res []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			res = append(res, Token{Plus, "+"})
		case '-':
			res = append(res, Token{Minus, "-"})
		case '(':
			res = append(res, Token{Lparen, "("})
		case ')':
			res = append(res, Token{Rparen, ")"})
		default:
			sb := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
					i++
				} else {
					res = append(res, Token{Int, sb.String()})
					i--
					break
				}
			}
		}
	}

	return res
}

// We can finally do our parsing! *ta-da*

// Recap:
// -> This demonstration showed how to basically implement the Interpreter pattern
// -> The idea is that typically we split it into two parts: Lexer and Parser
// -> Lexer: takes our textual input and split it up into a bunch of tokens
// -> Parser: takes those tokens and creates tree-like structures out of those tokens
// -> Those structures can be traversed, those that look like trees,
//	  in all sorts of ways

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)

	parsed := Parse(tokens)
	fmt.Printf("%s = %d\n", input, parsed.Value())
}
