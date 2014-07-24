package tree

import "github.com/ghigt/gocyk/grammar"

type Tree struct {
	Right *Tree
	Value grammar.Token
	Left  *Tree
}

func (t *Tree) InsertRight(v grammar.Token) {
	t.Right = New(v)
}

func (t *Tree) InsertLeft(v grammar.Token) {
	t.Left = New(v)
}

func New(v grammar.Token) *Tree {
	return &Tree{nil, v, nil}
}
