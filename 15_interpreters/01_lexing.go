// Interpreters - Lexing

// Aight, what we're going to do here for our demo is
// we're going to parse a simple numeric expression and
// then try to evaluate this.

// And this is going to be split into two parts,
// so the first part is this example.

// -> Lexing: taking inputs and spliting them up into separate tokens

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Now, the first question is -> what are those tokens?
// Where do we get them?

// Let's actually define a type for tokens.

type TokenType int

// <- We'll have a limited number of tokens

// Let's define a bunch of content, that we'll use here.

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

// With this, we can define a Lexing process.

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
		default: // not so obvious case for numbers
			// numbers can take up more than 1 character, and we need to process that
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

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)
}
