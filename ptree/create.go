package ptree

import (
	grm "github.com/ghigt/gocyk/grammar"
	"github.com/ghigt/gocyk/rtable"
)

func checkRight(rt *rtable.RTable, col int, row int, tok grm.Token) bool {
	itm := rt.GetItem(col, row)

	for _, t := range itm.GetTokens() {
		if t == tok {
			return true
		}
	}
	return false
}

//func buildTree(tst *rtable.TableST, tok grm.Token, row int, col int) *PTree {
//	for _, nt := range tst.Grammar.GetListOfNT(tok) {
//		left := nt.GetLeft()
//		right := nt.GetRight()
//		for c := col - 1; c >= row; c-- {
//			if itm := tst.GetItem(c, row); itm.isEmpty() != true {
//				for _, t := range itm.GetTokens() {
//					if t == left && checkRight(tst.Table, col, c+1, right) {
//						pt := New(tok)
//						pt.InsertLeft(left)
//						pt.InsertRight(right)
//						return pt
//					}
//				}
//			}
//		}
//	}
//	return nil
//}
//
//func Build(tst *rtable.TableST) []*PTree {
//	size := len(*tst)
//	pts := []*PTree{}
//
//	// iterate over rows
//	for row := 0; row < size; {
//		col := size - 1
//		for ; col >= row; col-- {
//			itm := tst.GetItem(col, row)
//			if itm.IsEmpty() == true {
//				if col == row {
//					// Add Terminals
//					//tst.Grammar.GetListOfT()
//				} else {
//					for _, tok := range itm.GetTokens() {
//						pts = append(pts, buildTree(tst, tok, row, col))
//					}
//				}
//				break
//			}
//		}
//		row = col + 1
//	}
//	return pts
//}
