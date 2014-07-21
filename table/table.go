package table

import (
	"bytes"

	gram "github.com/ghigt/gocyk/grammar"
)

// RTable is the type to build a recognition table.
type RTable []*Column

// grammar is the variable used in the grammar package.
var grammar *gram.Grammar

// New returns a new recognition table for the given grammar.
func New(g *gram.Grammar) *RTable {
	grammar = g

	return &RTable{}
}

// GetColumn returns the column corresponding to the given parameter.
func (rt *RTable) GetColumn(col int) *Column {
	return (*rt)[col]
}

// GetItem returns item corresponding to the column and index given
// in parameter.
func (rt *RTable) GetItem(column, index int) *Item {
	c := (*rt).GetColumn(column)

	return c.GetItem(index)
}

// Valid returns true if the recognition table is valid. This method
// checks if the top left item is empty or not.
func (rt *RTable) Valid() bool {
	if rt.GetItem(len(*rt)-1, 0).Empty() != true {
		return true
	}
	return false
}

// ValidFor returns true if the recognition table is valid for a given
// range. This method checks if the top left item is empty or not.
func (rt *RTable) ValidFor(beg, end int) bool {
	length := len(*rt)

	if beg < 0 || beg >= length ||
		end < 0 || end >= length ||
		end < beg {
		return false
	}
	if rt.GetItem(end, beg).Empty() != true {
		return true
	}
	return false
}

func (rt *RTable) String() string {
	var buf bytes.Buffer

	if rt == nil {
		return "nil"
	}
	for _, column := range *rt {
		buf.WriteString(column.String() + "\n")
	}

	return buf.String()
}
