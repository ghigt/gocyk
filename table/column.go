package table

import "bytes"

type Column []*Item

func (c *Column) GetItem(index int) *Item {
	return (*c)[index]
}

func (c *Column) AddAndCompute(s string, rt *RTable) {
	*c = make([]*Item, len(*rt))
	pos := len(*rt) - 1 // current position of the column

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

func (c *Column) String() string {
	var buf bytes.Buffer

	for _, item := range *c {
		buf.WriteString(item.String() + " ")
	}

	return buf.String()
}
