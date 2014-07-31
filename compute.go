package gocyk

import "github.com/ghigt/gocyk/rtable"

// CompleteColumn fills the given column with the correct items depending
// on the string, the position and the grammar.
func (g *GoCYK) CompleteColumn(s string, c *rtable.Column, pos int) {
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

func (g *GoCYK) CompleteColumnFrom(pos, col int) {
	c := g.Table.GetColumn(col)
	for i := pos; i >= 0; i-- {
		(*c)[i] = &rtable.Item{}
		for l := i; l < col; l++ {
			(*c)[i] = (*c)[i].Add(g.Grammar.GetTokensOfNT(
				g.Table.GetItem(l, i).GetTokens(),
				g.Table.GetItem(col, l+1).GetTokens(),
			)...)
		}
	}
}

func (g *GoCYK) CompleteFollowing(pos int) {
	for col := pos + 1; col < g.Table.Size(); col++ {
		(*g.Table)[col].AddFront(&rtable.Item{})
		g.CompleteColumnFrom(pos, col)
	}
}
