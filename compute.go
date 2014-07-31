package gocyk

import "github.com/ghigt/gocyk/rtable"

func (g *GoCYK) CompleteColumn(s string, c *rtable.Column, pos int) {
	for i := pos; i >= 0; i-- {
		if i == pos {
			(*c)[i] = (*c)[i].Add(g.Grammar.GetTokensOfT(s)...)
		} else {
			for l := i; l < pos; l++ {
				(*c)[i] = (*c)[i].Add(g.Grammar.GetTokensOfNT(
					(*g.Table).GetItem(l, i).GetTokens(),
					(*g.Table).GetItem(pos, l+1).GetTokens(),
				)...)
			}
		}
	}
}
