package rtable

import (
	"bytes"

	grm "github.com/ghigt/gocyk/grammar"
)

// Item type represents a slice of tokens for the recognition table.
type Item []grm.Token

// Create instantiate a new item if it is nil.
func (item *Item) Create() *Item {
	if item == nil {
		return new(Item)
	}
	return item
}

// IsPresent checks if the token is already present in the item.
func (item *Item) IsPresent(t grm.Token) bool {
	for _, i := range *item {
		if i == t {
			return true
		}
	}
	return false
}

// Add adds multiple tokens at the end of the item.
func (item *Item) Add(t grm.Token) *Item {
	if item == nil {
		item = new(Item)
	}
	if !item.IsPresent(t) {
		*item = append(*item, t)
	}
	return item
}

// GetTokens returns all the tokens of the item.
func (item *Item) GetTokens() []grm.Token {
	return *item
}

// IsEmpty checks if the item is empty or not.
func (item *Item) IsEmpty() bool {
	if item.Size() == 0 {
		return true
	}
	return false
}

// Size return the size the item.
func (item *Item) Size() int {
	return len(*item)
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
