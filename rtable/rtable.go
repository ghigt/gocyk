/*
  Package rtable manages and create a recognition table following
  the rules of the CYK (Cocke Younger Kasami) algorithm with a
  CNF (Chomsky Normal Form) grammar.
*/
package rtable

import (
	"bytes"
	"errors"
	"fmt"
)

// RTable is the type of the recognition table.
type RTable []*Column

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

// Add adds a new column at the end of the recognition table
func (rt *RTable) Add() *Column {
	c := NewColumn(rt.Size() + 1)
	*rt = append(*rt, c)
	return c
}

// Insert inserts a new column at give position of the recognition table.
func (rt *RTable) Insert(pos int) (*Column, error) {
	if pos < 0 || pos > rt.Size() {
		return nil,
			errors.New(fmt.Sprintf("index of table (%d) out of range", pos))
	}
	if pos == rt.Size() {
		return rt.Add(), nil
	}
	c := NewColumn(pos + 1)

	*rt = append(*rt, nil)
	copy((*rt)[pos+1:], (*rt)[pos:])
	(*rt)[pos] = c
	return c, nil
}

// Remove removes a column at give position of the recognition table.
func (rt *RTable) Remove(pos int) error {
	if pos < 0 || pos >= rt.Size() {
		return errors.New(fmt.Sprintf("index of table (%d) out of range", pos))
	}
	(*rt) = append((*rt)[:pos], (*rt)[pos+1:]...)
	return nil
}

func (rt *RTable) Size() int {
	return len(*rt)
}

// IsValid returns true if the recognition table is valid. This method
// checks if the top right item is empty or not.
func (rt *RTable) IsValid() bool {
	if rt.GetItem(rt.Size()-1, 0).IsEmpty() != true {
		return true
	}
	return false
}

// ValidFor returns true if the recognition table is valid for a given
// range. This method checks if the top right item is empty or not.
func (rt *RTable) IsValidFor(beg, end int) bool {
	length := rt.Size()

	if beg < 0 || beg >= length ||
		end < 0 || end >= length ||
		end < beg {
		return false
	}
	if rt.GetItem(end, beg).IsEmpty() != true {
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
