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
	return g.Insert(s, len(*g.Table))
}

func (g *GoCYK) Insert(s string, pos int) error {
	var c *rtable.Column

	c, err := g.Table.Insert(pos)
	if err != nil {
		return err
	}
	g.CompleteColumn(s, c, pos)
	return nil
}

func (g *GoCYK) Remove(s string, pos int) error {
	return nil
}

func (g *GoCYK) IsValid() bool {
	return g.Table.IsValid()
}
