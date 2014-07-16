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
	//GetToken() (Token, bool)
	//GetLeft() (Token, bool)
	//GetRight() (Token, bool)
}

type Rules []Rule

type Grammar map[Token]Rules

func (g *Grammar) GetTokensOfT(s string) (tokens []Token) {
	for t, rules := range *g {
		for _, r := range rules {
			v, ok := r.(string) // Terminal declared as type string
			if ok && string(v) == s {
				tokens = append(tokens, t)
			}
		}
	}
	return
}

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

func (g *Grammar) GetListOfT(t Token) (nt []*Terminal) {
	for _, i := range (*g)[t] {
		v, ok := i.(Terminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return
}

func (g *Grammar) GetListOfNT(t Token) (nt []*NonTerminal) {
	for _, i := range (*g)[t] {
		v, ok := i.(NonTerminal)
		if ok {
			nt = append(nt, &v)
		}
	}
	return
}
