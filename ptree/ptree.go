/*
  Package ptree manipulates the parsing tree after generation
  of a recognition table.
*/
package ptree

import "github.com/ghigt/gocyk/grammar"

type PTree struct {
	Right *PTree
	Value grammar.Token
	Left  *PTree
}

func (t *PTree) InsertRight(v grammar.Token) {
	t.Right = New(v)
}

func (t *PTree) InsertLeft(v grammar.Token) {
	t.Left = New(v)
}

func New(v grammar.Token) *PTree {
	return &Tree{nil, v, nil}
}
