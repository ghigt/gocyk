package table

type Column []*Item

func (c *Column) GetItem(index int) *Item {
	return (*c)[index], nil
}
