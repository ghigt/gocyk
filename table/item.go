package table

import (
	"bytes"

	gram "github.com/ghigt/gocyk/grammar"
)

type Item []gram.Token

func (item *Item) Add(t ...gram.Token) *Item {
	if item == nil {
		item = new(Item)
	}
	*item = append(*item, t...)
	return item
}

func (item *Item) GetTokens() []gram.Token {
	return *item
}

func (item *Item) Empty() bool {
	if len(*item) == 0 {
		return true
	}
	return false
}

func (item *Item) String() string {
	var buf bytes.Buffer

	if item == nil {
		return "nil"
	}
	buf.WriteString("{ ")
	for _, i := range *item {
		buf.WriteString(string(i) + " ")
	}
	buf.WriteString("}")

	return buf.String()
}
