package table

import "bytes"

type Column []*Item

func (c *Column) GetItem(index int) *Item {
	return (*c)[index]
}

func (c *Column) AddAndCompute(s string, rt *RTable) {
	*c = make([]*Item, len(*rt))

	for i := len(*rt) - 1; i >= 0; i-- {
		if i == len(*rt)-1 {
			(*c)[i] = (*c)[i].Add(grammar.GetTerminalTokens(s)...)
		} else {
			// Compute the end of the column
		}
	}
}

func (c *Column) String() string {
	var buf bytes.Buffer

	for _, item := range *c {
		buf.WriteString(item.String() + " ")
	}

	return buf.String()
}
