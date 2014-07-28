/*
  Package ptree manipulates the parsing tree after generation
  of a recognition table.
*/
package ptree

import (
	grm "github.com/ghigt/gocyk/grammar"
)

// PTree type is a structure which contains the right and left trees
// and also its own value.
type PTree struct {
	Right *PTree
	Value grm.Token
	Left  *PTree
}

// InsertRight insert and create a new tree for the right node.
func (t *PTree) InsertRight(v grm.Token) {
	t.Right = New(v)
}

// InsertLeft insert and create a new tree for the left node.
func (t *PTree) InsertLeft(v grm.Token) {
	t.Left = New(v)
}

// New returns a new tree with empty nodes and the given value.
func New(v grm.Token) *PTree {
	return &PTree{nil, v, nil}
}
