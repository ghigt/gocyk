package gocyk

import (
	grm "github.com/ghigt/gocyk/grammar"
	"github.com/ghigt/gocyk/ptree"
	"github.com/ghigt/gocyk/rtable"
)

type GoCYK struct {
	Table   *rtable.RTable
	Tree    *ptree.PTree
	Grammar *grm.Grammar
}

func New(grammar *grm.Grammar) *GoCYK {
	return &GoCYK{
		Table:   &rtable.RTable{},
		Tree:    &ptree.PTree{},
		Grammar: grammar,
	}
}

func (g *GoCYK) Add(s string) error {
	return g.Insert(s, g.Table.Size())
}

func (g *GoCYK) Insert(s string, pos int) error {
	c, err := g.Table.Insert(pos)
	if err != nil {
		return err
	}
	g.completeColumn(s, c, pos)
	g.insertFollowing(pos)
	return nil
}

func (g *GoCYK) Remove(pos int) error {
	if err := g.Table.Remove(pos); err != nil {
		return err
	}
	g.removeFollowing(pos)
	return nil
}

func (g *GoCYK) IsValid() bool {
	return g.Table.IsValid()
}
