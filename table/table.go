package table

import (
	"bytes"
	"errors"
)

// RTable is the type to build a recognition table
type RTable []*Column

func (rt *RTable) GetColumn(index int) (*Column, error) {
	if rt == nil {
		return nil, errors.New("Error: null pointer")
	}
	return (*rt)[index], nil
}

func (rt *RTable) Add(s string) error {
	if rt == nil {
		return errors.New("Error: null pointer")
	}
	//lengthT := make([]*Item, len(*rt)+1)

	// -- TEST --
	// -- TEST --

	return nil
}

func (rt *RTable) Insert(index int) error {
	if rt == nil {
		return errors.New("Error: null pointer")
	}
	return nil
}

func (rt *RTable) Remove(index int) error {
	if rt == nil {
		return errors.New("Error: null pointer")
	}
	return nil
}

func (rt *RTable) Valid() bool {
	if rt == nil {
		return false
	}
	return false
}

func (rt *RTable) ValidFor(beg, end int) bool {
	if rt == nil {
		return false
	}
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
