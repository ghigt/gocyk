package table

import "errors"

type Column []*Item

func (c *Column) GetItem(index int) (*Item, error) {
	if c == nil {
		return nil, errors.New("Error: null pointer")
	}
	return (*c)[index], nil
}
