package table

type Column []*Item

func (c *Column) GetItem(index int) *Item {
	return (*c)[index], nil
}

func (c *Column) AddAndCompute(s string, rt *RTable) {
	*c = make([]*Item, len(*rt))

	for i := len(*rt) - 1; i >= 0; i-- {
		if i == len(*rt)-1 {
			c[i].Add(grammar.GetTerminalTokens(s))
		}
	}
}
