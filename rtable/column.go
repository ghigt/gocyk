package rtable

import (
	"bytes"
	"errors"
	"fmt"
)

// Column type manages each column of the table package
type Column []*Item

// GetItem returns the *Item present at the `index`.
func (c *Column) GetItem(index int) *Item {
	return (*c)[index]
}

// AddFront adds a new item at the front of the column.
func (c *Column) AddFront(item *Item) {
	*c = append(*c, nil)
	copy((*c)[1:], (*c)[:])
	(*c)[0] = item
}

// Remove removes the first item of the column.
func (c *Column) Remove(pos int) error {
	if pos < 0 || pos >= len(*c) {
		return errors.New(fmt.Sprintf("index of column (%d) out of range",
			pos))
	}
	(*c) = append((*c)[:pos], (*c)[pos+1:]...)
	return nil
}

func (c *Column) String() string {
	var buf bytes.Buffer

	for _, item := range *c {
		buf.WriteString(item.String() + " ")
	}
	return buf.String()
}
