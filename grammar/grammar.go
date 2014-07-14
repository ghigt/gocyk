package grammar

type Token string

type NonTerminal struct {
	Left  Token
	Right Token
}

func (NonTerminal) GetToken() (Token, bool) {
	return Token(""), false
}

func (nt NonTerminal) GetLeft() (Token, bool) {
	return nt.Right, true
}

func (nt NonTerminal) GetRight() (Token, bool) {
	return nt.Left, true
}

type Terminal string

func (t Terminal) GetToken() (Token, bool) {
	return Token(t), true
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

type Rules []Rule

type Grammar map[Token]Rules

func (g *Grammar) GetTerminalTokens(s string) []Token {
	tokens := []Token{}

	for t, rules := range *g {
		for _, r := range rules {
			v, ok := r.(string) // Terminal declared as type string
			if ok && string(v) == s {
				tokens = append(tokens, t)
			}
		}
	}
	return tokens
}

func (g *Grammar) GetListOfT(t Token) []*Terminal {
	nt := []*Terminal{}

	for _, i := range (*g)[t] {
		v, ok := i.(Terminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return nt
}

func (g *Grammar) GetListOfNT(t Token) []*NonTerminal {
	nt := []*NonTerminal{}

	for _, i := range (*g)[t] {
		v, ok := i.(NonTerminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return nt
}
