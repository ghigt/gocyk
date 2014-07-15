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

	return &RTable{}
}

func (rt *RTable) GetColumn(index int) *Column {
	return (*rt)[index]
}

func (rt *RTable) GetItem(column, index int) *Item {
	c := (*rt).GetColumn(column)

	return c.GetItem(index)
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
	if rt.GetItem(len(*rt)-1, 0).Empty() != true {
		return true
	}
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
	for _, column := range *rt {
		buf.WriteString(column.String() + "\n")
	}

	return buf.String()
}
