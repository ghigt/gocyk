package table

import "errors"

func (rt *RTable) Add(s string) {
	c := new(Column)
	*rt = append(*rt, c)
	c.AddAndCompute(s, rt)
}

func (rt *RTable) insertFollowing(pos int) {
	for col := pos + 1; col < len(*rt); col++ {
		(*rt)[col].Front(&Item{})
		(*rt)[col].ComputeFrom(pos, col, rt)
	}
}

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

func (rt *RTable) removeFollowing(pos int) {
	for col := pos; col < len(*rt); col++ {
		(*rt)[col].PopFront()
		(*rt)[col].ComputeFrom(pos-1, col, rt)
	}
}

func (rt *RTable) Remove(pos int) error {
	if pos < 0 || pos >= len(*rt) {
		return errors.New("index out of range")
	}
	(*rt) = append((*rt)[:pos], (*rt)[pos+1:]...)
	rt.removeFollowing(pos)
	return nil
}
