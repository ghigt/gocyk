package rtable

import "bytes"

// Column type manages each column of the table package
type Column []*Item

// GetItem returns the *Item present at the `index`.
func (c *Column) GetItem(index int) *Item {
	return (*c)[index]
}

// Front adds a new `item` at the front of the column.
func (c *Column) Front(item *Item) {
	*c = append(*c, nil)
	copy((*c)[1:], (*c)[:])
	(*c)[0] = item
}

// PopFront remove the first *Item of the column.
func (c *Column) PopFront() {
	*c = (*c)[1:]
}

func (c *Column) String() string {
	var buf bytes.Buffer

	for _, item := range *c {
		buf.WriteString(item.String() + " ")
	}
	return buf.String()
}
