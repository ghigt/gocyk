package rtable

// insertFollowing computes the rest each item until the end and
// modify it appropriately. It also modifies only from a given position.
func (rt *RTable) insertFollowing(pos int) {
	for col := pos + 1; col < len(*rt); col++ {
		(*rt)[col].Front(&Item{})
		(*rt)[col].ComputeFrom(pos, col, rt)
	}
}

// removeFollowing computes the rest each item until the end and
// modify it appropriately. It also modifies only from a given position.
func (rt *RTable) removeFollowing(pos int) {
	for col := pos; col < len(*rt); col++ {
		(*rt)[col].PopFront()
		(*rt)[col].ComputeFrom(pos-1, col, rt)
	}
}
