package grammar

import "bytes"

type Token string

type NonTerminal struct {
	Left  Token
	Right Token
}

type Terminal string

type Rule interface{}

type Items []Token

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
