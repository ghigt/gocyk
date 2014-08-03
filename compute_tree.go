package gocyk

import (
	"log"

	grm "github.com/ghigt/gocyk/grammar"
	"github.com/ghigt/gocyk/ptree"
)

func (g *GoCYK) checkRight(col int, row int, tok grm.Token) bool {
	itm := g.Table.GetItem(col, row)

	for _, t := range itm.GetTokens() {
		if t == tok {
			return true
		}
	}
	return false
}

func (g *GoCYK) buildTree(tok grm.Token, row int, col int) *ptree.PTree {
	pt := ptree.New(tok)
	if col == row {
		return pt
	} else {
		for _, nt := range g.Grammar.GetListOfNT(tok) {
			left, _ := nt.GetLeft()
			right, _ := nt.GetRight()
			for c := col - 1; c >= row; c-- {
				if itm := g.Table.GetItem(c, row); itm.IsEmpty() != true {
					for _, t := range itm.GetTokens() {
						if t == left && g.checkRight(col, c+1, right) {
							pt.Left = g.buildTree(left, row, c)
							pt.Right = g.buildTree(right, c+1, col)
							return pt
						}
					}
				}
			}
		}
	}
	return nil
}

func (g *GoCYK) BuildTrees() []*ptree.PTree {
	size := g.Table.Size()
	pts := []*ptree.PTree{}

	for row := 0; row < size; {
		col := size - 1
		for ; col >= row; col-- {
			itm := g.Table.GetItem(col, row)
			if itm.IsEmpty() == false {
				for _, tok := range itm.GetTokens() {
					if tree := g.buildTree(tok, row, col); tree != nil {
						pts = append(pts, tree)
					} else {
						log.Fatal("nil Tree")
					}
				}
				break
			}
		}
		row = col + 1
	}
	return pts
}
