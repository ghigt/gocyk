/*
  Package ptree manipulates the parsing tree after generation
  of a recognition table.
*/
package ptree

import "github.com/ghigt/gocyk/grammar"

// PTree type is a structure which contains the right and left trees
// and also its own value.
type PTree struct {
	Right *PTree
	Value grammar.Token
	Left  *PTree
}

// InsertRight insert and create a new tree for the right node.
func (t *PTree) InsertRight(v grammar.Token) {
	t.Right = New(v)
}

// InsertLeft insert and create a new tree for the left node.
func (t *PTree) InsertLeft(v grammar.Token) {
	t.Left = New(v)
}

// New returns a new tree with empty nodes and the given value.
func New(v grammar.Token) *PTree {
	return &Tree{nil, v, nil}
}
