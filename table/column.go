package table

import "bytes"

type Column []*Item

func (c *Column) GetItem(index int) *Item {
	return (*c)[index]
}

func (c *Column) Front(item *Item) {
	*c = append(*c, nil)
	copy((*c)[1:], (*c)[:])
	(*c)[0] = item
}

func (c *Column) PopFront() {
	*c = (*c)[1:]
}

func (c *Column) compute(s string, pos int, rt *RTable) {
	for i := pos; i >= 0; i-- {
		if i == pos {
			(*c)[i] = (*c)[i].Add(grammar.GetTokensOfT(s)...)
		} else {
			for l := i; l < pos; l++ {
				(*c)[i] = (*c)[i].Add(grammar.GetTokensOfNT(
					(*rt).GetItem(l, i).GetTokens(),
					(*rt).GetItem(pos, l+1).GetTokens(),
				)...)
			}
		}
	}
}

func (c *Column) ComputeFrom(pos int, col int, rt *RTable) {
	for i := pos; i >= 0; i-- {
		(*c)[i] = &Item{}
		for l := i; l < col; l++ {
			(*c)[i] = (*c)[i].Add(grammar.GetTokensOfNT(
				(*rt).GetItem(l, i).GetTokens(),
				(*rt).GetItem(col, l+1).GetTokens(),
			)...)
		}
	}
}

func (c *Column) AddAndCompute(s string, rt *RTable) {
	*c = make([]*Item, len(*rt))
	pos := len(*rt) - 1 // current position of the column

	c.compute(s, pos, rt)
}

func (c *Column) InsertAndCompute(s string, pos int, rt *RTable) {
	*c = make([]*Item, pos+1)

	c.compute(s, pos, rt)
}

func (c *Column) String() string {
	var buf bytes.Buffer

	for _, item := range *c {
		buf.WriteString(item.String() + " ")
	}

	return buf.String()
}
