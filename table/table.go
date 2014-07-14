package table

import (
	"bytes"

	gram "github.com/ghigt/gocyk/grammar"
)

// RTable is the type to build a recognition table
type RTable []*Column

var grammar *gram.Grammar

func New(g *gram.Grammar) *RTable {
	grammar = g

	return &RTable
}

func (rt *RTable) GetColumn(index int) *Column {
	return (*rt)[index]
}

func (rt *RTable) GetItem(column, index, int) *Item {
	column := (*rt).GetColumn(column)

	return column.GetItem(index)
}

func (rt *RTable) Add(s string) {
	c := new(Column)
	*rt = append(*rt, c)
	c.AddAndCompute(s, rt)
}

func (rt *RTable) Insert(index int) error {
	return nil
}

func (rt *RTable) Remove(index int) error {
	return nil
}

func (rt *RTable) Valid() bool {
	return false
}

func (rt *RTable) ValidFor(beg, end int) bool {
	return false
}

func (rt *RTable) String() string {
	var buf bytes.Buffer

	if rt == nil {
		return "nil"
	}
	buf.WriteString("[length][index]\n")
	//for i, index := range *rt {
	// Add column
	//buf.WriteString(index[l].String() + " ")
	//}

	return buf.String()
}
