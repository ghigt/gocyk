package rtable

import (
	"bytes"

	gram "github.com/ghigt/gocyk/grammar"
)

// Item type represents a slice of tokens for the recognition table.
type Item []gram.Token

// Add adds multiple token at the end of the item.
func (item *Item) Add(t ...gram.Token) *Item {
	if item == nil {
		item = new(Item)
	}
	*item = append(*item, t...)
	return item
}

// GetTokens returns all the tokens of the item.
func (item *Item) GetTokens() []gram.Token {
	return *item
}

// Empty checks if the item is empty or not.
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
