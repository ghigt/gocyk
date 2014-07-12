package grammar

import "bytes"

type Token string

type NonTerminal struct {
	Left  Token
	Right Token
}

func (NonTerminal) GetTerminal() (Terminal, bool) {
	return Terminal(""), false
}

func (nt NonTerminal) GetLeft() (Token, bool) {
	return nt.Right, true
}

func (nt NonTerminal) GetRight() (Token, bool) {
	return nt.Left, true
}

type Terminal string

func (t Terminal) GetTerminal() (Terminal, bool) {
	return t, true
}

func (Terminal) GetLeft() (Token, bool) {
	return Token(""), false
}

func (Terminal) GetRight() (Token, bool) {
	return Token(""), false
}

type Rule interface {
	//GetTerminal() (Terminal, bool)
	//GetLeft() (Token, bool)
	//GetRight() (Token, bool)
}

type Items []Token

func (items *Items) Add(t Token) *Items {
	if items == nil {
		items = new(Items)
	}
	*items = append(*items, t)
	return items
}

func (items *Items) String() string {
	var buf bytes.Buffer

	buf.WriteString("{ ")
	for _, i := range *items {
		buf.WriteString(string(i) + " ")
	}
	buf.WriteString("}")

	return buf.String()
}

type Rules []Rule

type Grammar map[Token]Rules

func (g Grammar) GetListOfT(t Token) []*Terminal {
	nt := []*Terminal{}

	for _, i := range g[t] {
		v, ok := i.(Terminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return nt
}

func (g Grammar) GetListOfNT(t Token) []*NonTerminal {
	nt := []*NonTerminal{}

	for _, i := range g[t] {
		v, ok := i.(NonTerminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return nt
}
