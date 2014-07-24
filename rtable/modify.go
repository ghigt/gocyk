package rtable

import "errors"

// Add adds a new column at the end of the recognition table and
// fills it with the appropriate items.
func (rt *RTable) Add(s string) {
	c := new(Column)
	*rt = append(*rt, c)
	c.AddAndCompute(s, rt)
}

// insertFollowing computes the rest each item until the end and
// modify it appropriately. It also modifies only from a given position.
func (rt *RTable) insertFollowing(pos int) {
	for col := pos + 1; col < len(*rt); col++ {
		(*rt)[col].Front(&Item{})
		(*rt)[col].ComputeFrom(pos, col, rt)
	}
}

// Insert inserts a new column at give position of the
// recognition table and fills it with the appropriate items.
func (rt *RTable) Insert(s string, pos int) error {
	if pos < 0 || pos > len(*rt) {
		return errors.New("index out of range")
	}
	if pos == len(*rt) {
		rt.Add(s)
		return nil
	}
	c := new(Column)

	*rt = append(*rt, nil)
	copy((*rt)[pos+1:], (*rt)[pos:])
	(*rt)[pos] = c
	c.InsertAndCompute(s, pos, rt)

	rt.insertFollowing(pos)
	return nil
}

// removeFollowing computes the rest each item until the end and
// modify it appropriately. It also modifies only from a given position.
func (rt *RTable) removeFollowing(pos int) {
	for col := pos; col < len(*rt); col++ {
		(*rt)[col].PopFront()
		(*rt)[col].ComputeFrom(pos-1, col, rt)
	}
}

// Remove removes a column at give position of the
// recognition table and compute the new items appropriately.
func (rt *RTable) Remove(pos int) error {
	if pos < 0 || pos >= len(*rt) {
		return errors.New("index out of range")
	}
	(*rt) = append((*rt)[:pos], (*rt)[pos+1:]...)
	rt.removeFollowing(pos)
	return nil
}
