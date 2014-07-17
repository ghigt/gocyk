package table

import (
	"bytes"
	"errors"

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

func (rt *RTable) modifyFollowing(pos int) {
	for col := pos + 1; col < len(*rt); col++ {
		(*rt)[col].Front(&Item{})
		(*rt)[col].ComputeFrom(pos, col, rt)
	}
}

func (rt *RTable) Insert(s string, pos int) error {
	if pos < 0 || pos > len(*rt) {
		return errors.New("index out of range")
	}
	if pos == len(*rt) {
		rt.Add(s)
		return nil
	}
	c := new(Column)

	*rt = append(*rt, nil)
	copy((*rt)[pos+1:], (*rt)[pos:])
	(*rt)[pos] = c
	c.InsertAndCompute(s, pos, rt)

	rt.modifyFollowing(pos)
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
