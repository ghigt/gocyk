// Package grammar implements utility for manipulating
// and create a CNF (Chomsky Normal Form) grammar.
package grammar

import (
	"fmt"
	"regexp"
)

// Token type represent the symbol or a substring of
// a grammar.
type Token string

// NonTerminal is the type respresenting the right-hand side
// of a rule in CNF with 2 symbols.
type NonTerminal struct {
	Left  Token
	Right Token
}

// GetToken returns the token of a Rule. A NonTerminal doesn't
// have the token.
func (NonTerminal) GetToken() (Token, bool) {
	return Token(""), false
}

// GetLeft returns the left symbol of the right-hand side rule.
func (nt NonTerminal) GetLeft() (Token, bool) {
	return nt.Left, true
}

// GetRight returns the right symbol of the right-hand side rule.
func (nt NonTerminal) GetRight() (Token, bool) {
	return nt.Right, true
}

// Terminal is the type representing the right-hand side
// of a rule in CNF with only one terminal.
type Terminal string

// GetToken returns the Token of the Rule.
func (t Terminal) GetToken() (Token, bool) {
	return Token(t), true
}

// GetLeft returns the left symbol of the Rule. For a Terminal,
// No Token is returned.
func (Terminal) GetLeft() (Token, bool) {
	return Token(""), false
}

// GetRight returns the right symbol of the Rule. For a Terminal,
// No Token is returned.
func (Terminal) GetRight() (Token, bool) {
	return Token(""), false
}

// Rule is the interface representing the right-hand side of
// a CNF grammar.
type Rule interface {
	//GetToken() (Token, bool)
	//GetLeft() (Token, bool)
	//GetRight() (Token, bool)
}

// Rules is a slice of Rule representing the different possibilities
// of each symbol.
type Rules []Rule

// Grammar is the type representing the CNF grammar.
type Grammar map[Token]Rules

// GetTokensOfT returns all the Terminal tokens matching the
// right-hand side given in paramater `s`.
func (g *Grammar) GetTokensOfT(s string) (tokens []Token) {
	for t, rules := range *g {
		for _, r := range rules {
			v, ok := r.(string) // Terminal declared as type string
			if ok && IsRegEq(v, s) {
				tokens = append(tokens, t)
			}
		}
	}
	return
}

// GetTokensOfNT returns all the NonTerminal tokens matching the
// right-hand side given in paramater.
func (g *Grammar) GetTokensOfNT(leftT []Token, rightT []Token) (tokens []Token) {
	for t, rules := range *g {
		for _, r := range rules {
			v, ok := r.(NonTerminal)
			if ok {
				for _, left := range leftT {
					if left == v.Left {
						for _, right := range rightT {
							if right == v.Right {
								tokens = append(tokens, t)
							}
						}
					}
				}
			}
		}
	}
	return
}

// GetListOfT returns a list of *Terminal matching with the
// left-hand side token `t` given in parameter.
func (g *Grammar) GetListOfT(t Token) (nt []*Terminal) {
	for _, i := range (*g)[t] {
		v, ok := i.(Terminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return
}

// GetListOfNT returns a list of *NonTerminal matching with the
// left-hand side token `t` given in parameter.
func (g *Grammar) GetListOfNT(t Token) (nt []*NonTerminal) {
	for _, i := range (*g)[t] {
		v, ok := i.(NonTerminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return
}

// IsRegEq check if the `reg` match as a regular expression with
// the value `val`.
func IsRegEq(reg string, val string) bool {
	exp := fmt.Sprintf("^%s$", reg)
	res, _ := regexp.MatchString(exp, val)
	fmt.Printf("match %q with %s: %v\n", exp, val, res)

	return res
}
