package gocyk

import "github.com/ghigt/gocyk/rtable"

func (g *GoCYK) CompleteColumn(s string, c *rtable.Column, pos int) {
	for i := pos; i >= 0; i-- {
		if i == pos {
			item := c.GetItem(i)
			c.SetItem(item.Add(g.Grammar.GetTokensOfT(s)...), i)
		} else {
			for l := i; l < pos; l++ {
				item := c.GetItem(i)
				c.SetItem(item.Add(g.Grammar.GetTokensOfNT(
					(*g.Table).GetItem(l, i).GetTokens(),
					(*g.Table).GetItem(pos, l+1).GetTokens(),
				)...), i)
			}
		}
	}
}
