package gocyk

import "github.com/ghigt/gocyk/rtable"

// CompleteColumn fills the given column with the correct items depending
// on the string, the position and the grammar.
func (g *GoCYK) completeColumn(s string, c *rtable.Column, pos int) {
	for i := pos; i >= 0; i-- {
		if i == pos {
			item := c.GetItem(i)
			c.SetItem(item.Add(g.Grammar.GetTokensOfT(s)...), i)
		} else {
			for l := i; l < pos; l++ {
				item := c.GetItem(i)
				c.SetItem(item.Add(g.Grammar.GetTokensOfNT(
					g.Table.GetItem(l, i).GetTokens(),
					g.Table.GetItem(pos, l+1).GetTokens(),
				)...), i)
			}
		}
	}
}

// completeColumnFrom recompute the column from a given position of index.
func (g *GoCYK) completeColumnFrom(pos, col int) {
	c := g.Table.GetColumn(col)
	for i := pos; i >= 0; i-- {
		c.SetItem(&rtable.Item{}, i)
		for l := i; l < col; l++ {
			item := c.GetItem(i)
			c.SetItem(item.Add(g.Grammar.GetTokensOfNT(
				g.Table.GetItem(l, i).GetTokens(),
				g.Table.GetItem(col, l+1).GetTokens(),
			)...), i)
		}
	}
}

// insertFollowing add a new item at the front of the right-hand of the
// position and recompute the items in the recognition table from
// the given position until the end of the table.
func (g *GoCYK) insertFollowing(pos int) {
	for col := pos + 1; col < g.Table.Size(); col++ {
		g.Table.GetColumn(col).AddFront(&rtable.Item{})
		g.completeColumnFrom(pos, col)
	}
}

// removeFollowing removes the first item of the right-hand of the
// position and computes the rest each item until the end and modify
// it appropriately. It also modifies only from a given position.
func (g *GoCYK) removeFollowing(pos int) {
	for col := pos; col < g.Table.Size(); col++ {
		g.Table.GetColumn(col).Remove(0)
		g.completeColumnFrom(pos-1, col)
	}
}
