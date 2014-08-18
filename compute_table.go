package gocyk

import (
	"sync"

	"github.com/ghigt/gocyk/rtable"
)

// CompleteColumn fills the given column with the correct items depending
// on the string, the position and the grammar.
func (g *GoCYK) completeColumn(s string, c *rtable.Column, pos int) {
	for i := pos; i >= 0; i-- {
		if i == pos {
			item := c.GetItem(i)
			item = item.Create()
			for _, t := range g.Grammar.GetTokensOfT(s) {
				item = item.Add(t)
			}
			c.SetItem(item, i)
		} else {
			for l := i; l < pos; l++ {
				item := c.GetItem(i)
				item = item.Create()
				for _, t := range g.Grammar.GetTokensOfNT(
					g.Table.GetItem(l, i).GetTokens(),
					g.Table.GetItem(pos, l+1).GetTokens()) {

					item = item.Add(t)
				}
				c.SetItem(item, i)
			}
		}
	}
}

// completeColumnFrom recompute the column from a given position of index.
func (g *GoCYK) completeColumnFrom(left <-chan struct{}, right chan<- struct{}, pos, col int) {
	c := g.Table.GetColumn(col)
	for i := pos; i >= 0; i-- {
		if left != nil {
			<-left
		}
		c.SetItem(&rtable.Item{}, i)
		for l := i; l < col; l++ {
			item := c.GetItem(i)
			item = item.Create()
			for _, t := range g.Grammar.GetTokensOfNT(
				g.Table.GetItem(l, i).GetTokens(),
				g.Table.GetItem(col, l+1).GetTokens()) {

				item = item.Add(t)
			}
			c.SetItem(item, i)
		}
		right <- struct{}{}
	}
}

// insertFollowing add a new item at the front of the right-hand of the
// position and recompute the items in the recognition table from
// the given position until the end of the table.
func (g *GoCYK) insertFollowing(pos int) {
	var left chan struct{}
	var wg sync.WaitGroup

	for col := pos + 1; col < g.Table.Size(); col++ {
		wg.Add(1)
		right := make(chan struct{}, col+1)
		go func(left chan struct{}, right chan struct{}, col int, pos int) {
			defer wg.Done()
			g.Table.GetColumn(col).AddFront(&rtable.Item{})
			g.completeColumnFrom(left, right, pos, col)
		}(left, right, col, pos)
		left = right
	}
	wg.Wait()
}

// removeFollowing removes the first item of the right-hand of the
// position and computes the rest each item until the end and modify
// it appropriately. It also modifies only from a given position.
func (g *GoCYK) removeFollowing(pos int) {
	var left chan struct{}
	var wg sync.WaitGroup

	for col := pos; col < g.Table.Size(); col++ {
		wg.Add(1)
		right := make(chan struct{}, col+1)
		go func(left chan struct{}, right chan struct{}, col int, pos int) {
			defer wg.Done()
			g.Table.GetColumn(col).Remove(0)
			g.completeColumnFrom(left, right, pos-1, col)
		}(left, right, col, pos)
		left = right
	}
	wg.Wait()
}

// completeColumnFrom recompute the column from a given position of index
// (without concurrency).
func (g *GoCYK) completeColumnFromWithoutC(pos, col int) {
	c := g.Table.GetColumn(col)
	for i := pos; i >= 0; i-- {
		c.SetItem(&rtable.Item{}, i)
		for l := i; l < col; l++ {
			item := c.GetItem(i)
			item = item.Create()
			for _, t := range g.Grammar.GetTokensOfNT(
				g.Table.GetItem(l, i).GetTokens(),
				g.Table.GetItem(col, l+1).GetTokens()) {

				item = item.Add(t)
			}
			c.SetItem(item, i)
		}
	}
}

// insertFollowingWithoutC add a new item at the front of the right-hand of the
// position and recompute the items in the recognition table from
// the given position until the end of the table (without concurrency).
func (g *GoCYK) insertFollowingWithoutC(pos int) {

	for col := pos + 1; col < g.Table.Size(); col++ {
		g.Table.GetColumn(col).AddFront(&rtable.Item{})
		g.completeColumnFromWithoutC(pos, col)
	}
}

// removeFollowing removes the first item of the right-hand of the
// position and computes the rest each item until the end and modify
// it appropriately. It also modifies only from a given position
// (without concurrency).
func (g *GoCYK) removeFollowingWithoutC(pos int) {

	for col := pos; col < g.Table.Size(); col++ {
		g.Table.GetColumn(col).Remove(0)
		g.completeColumnFromWithoutC(pos-1, col)
	}
}
