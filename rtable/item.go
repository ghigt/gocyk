package rtable

import (
	"bytes"

	grm "github.com/ghigt/gocyk/grammar"
)

// Item type represents a slice of tokens for the recognition table.
type Item []grm.Token

// Add adds multiple tokens at the end of the item.
func (item *Item) Add(t ...grm.Token) *Item {
	if item == nil {
		item = new(Item)
	}
	*item = append(*item, t...)
	return item
}

// GetTokens returns all the tokens of the item.
func (item *Item) GetTokens() []grm.Token {
	return *item
}

// IsEmpty checks if the item is empty or not.
func (item *Item) IsEmpty() bool {
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
