package ptree

import (
	grm "github.com/ghigt/gocyk/grammar"
	"github.com/ghigt/gocyk/rtable"
)

func checkRight(rt *RTable, col int, row int, tok grm.Token) {
	itm := rt.GetItem(col, row)

	for _, t := range itm.GetTokens() {
		if t == tok {
			return true
		}
	}
	return false
}

func buildTree(rt *RTable, tok grm.Token, row int, col int) *PTree {
	pt := *Ptree{}

	for _, nt := range rt.Grammar.GetListOfNT(tok) {
		left := nt.GetLeft()
		right := nt.GetRight()
		for c := col - 1; c >= row; c-- {
			if itm := rt.GetItem(c, row); itm.isEmpty() != true {
				for _, t := range itm.GetTokens() {
					if t == left && checkRight(rt, col, c+1, right) {
						// GOOD
					}
				}
			}
		}
	}
}

func Build(rt *rtable.RTable) []*PTree {
	size := len(*rt)
	pts := []*PTree{}

	// iterate over rows
	for row := 0; row < size; {
		col := size - 1
		for ; col >= row; col-- {
			itm := rt.GetItem(col, row)
			if itm.IsEmpty() == true {
				if col == row {
					// Add Terminals
					//rt.Grammar.GetListOfT()
				} else {
					for _, tok := range itm.GetTokens() {
						pts = pts.append(pts, buildTree(rt, tok, row, col))
					}
				}
				break
			}
		}
		row = col + 1
	}
}
