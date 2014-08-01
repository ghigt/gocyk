package rtable

import (
	"bytes"
	"fmt"
)

// Column type manages each column of the table package
type Column []*Item

// NewColumn returns a *Column initialized with a given size of items
func NewColumn(size int) *Column {
	var c Column

	c = make([]*Item, size)
	return &c
}

// GetItem returns the *Item present at the given position.
func (c *Column) GetItem(pos int) *Item {
	return (*c)[pos]
}

// SetItem set the *Item into the column at the given position.
func (c *Column) SetItem(item *Item, pos int) {
	(*c)[pos] = item
}

// AddFront adds a new item at the front of the column.
func (c *Column) AddFront(item *Item) {
	*c = append(*c, nil)
	copy((*c)[1:], (*c)[:])
	(*c)[0] = item
}

// Remove removes the first item of the column.
func (c *Column) Remove(pos int) error {
	if pos < 0 || pos >= c.Size() {
		return fmt.Errorf("index of column (%d) out of range", pos)
	}
	(*c) = append((*c)[:pos], (*c)[pos+1:]...)
	return nil
}

// Size returns the size of the column.
func (c *Column) Size() int {
	return len(*c)
}

func (c *Column) String() string {
	var buf bytes.Buffer

	for _, item := range *c {
		buf.WriteString(item.String() + " ")
	}
	return buf.String()
}
