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

// compute fills the table at the position `pos` with the
// appropriate items corresponding to the token `s` given in parameter.
func (c *Column) compute(s string, pos int, tst *TableST) {
	for i := pos; i >= 0; i-- {
		if i == pos {
			(*c)[i] = (*c)[i].Add(tst.Grammar.GetTokensOfT(s)...)
		} else {
			for l := i; l < pos; l++ {
				(*c)[i] = (*c)[i].Add(tst.Grammar.GetTokensOfNT(
					(*tst).GetItem(l, i).GetTokens(),
					(*tst).GetItem(pos, l+1).GetTokens(),
				)...)
			}
		}
	}
}

// ComputeFrom fills the table at the position `pos` with the
// appropriate items frome a certain column `col` until the end of
// the table.
func (c *Column) ComputeFrom(pos int, col int, tst *TableST) {
	for i := pos; i >= 0; i-- {
		(*c)[i] = &Item{}
		for l := i; l < col; l++ {
			(*c)[i] = (*c)[i].Add(tst.Grammar.GetTokensOfNT(
				(*tst).GetItem(l, i).GetTokens(),
				(*tst).GetItem(col, l+1).GetTokens(),
			)...)
		}
	}
}

// AddAndCompute fills a fresh new column with the appropriate
// items corresponding to the token `s` given in parameter.
func (c *Column) AddAndCompute(s string, tst *TableST) {
	*c = make([]*Item, len(*tst.RTable))
	pos := len(*tst.RTable) - 1 // current position of the column

	c.compute(s, pos, tst)
}

// InsertAndCompute fills a fresh new column with the appropriate
// items corresponding to the token `s` at a certain position `pos`.
func (c *Column) InsertAndCompute(s string, pos int, tst *TableST) {
	*c = make([]*Item, pos+1)

	c.compute(s, pos, tst)
}

func (c *Column) String() string {
	var buf bytes.Buffer

	for _, item := range *c {
		buf.WriteString(item.String() + " ")
	}

	return buf.String()
}
