// Package gocyk represents an abstraction of the CYK
// (Cocke Younger Kasami) parsing algorithm using a CNF
// (Chomsky Norm Form) for the grammar.
package gocyk

import (
	grm "github.com/ghigt/gocyk/grammar"
	"github.com/ghigt/gocyk/ptree"
	"github.com/ghigt/gocyk/rtable"
)

// GoCYK type contains the recognition table, the parsing tree,
// the grammar and the substrings. It provides methods to abstract
// the calculation of the CYK algorithm to modify the parsing text.
type GoCYK struct {
	Table   *rtable.RTable
	Tree    *ptree.PTree
	Grammar *grm.Grammar
	Sub     []string
}

// New instanciate a GoCYK structure and returns its address.
func New(grammar *grm.Grammar) *GoCYK {
	return &GoCYK{
		Table:   &rtable.RTable{},
		Tree:    &ptree.PTree{},
		Grammar: grammar,
	}
}

// Add adds a new substring at the end then recalculate the
// recognition table and the parsing tree.
func (g *GoCYK) Add(s string) error {
	return g.Insert(s, g.Table.Size())
}

// Insert inserts a new substring at the given position then
// recalculate the recognition table and the parsing tree.
func (g *GoCYK) Insert(s string, pos int) error {
	c, err := g.Table.Insert(pos)
	if err != nil {
		return err
	}
	g.Sub = append(g.Sub, "")
	copy(g.Sub[pos+1:], g.Sub[pos:])
	g.Sub[pos] = s
	g.completeColumn(s, c, pos)
	if pos < g.Table.Size()-1 {
		g.insertFollowing(pos)
	}
	return nil
}

// InsertNC inserts a new substring at the given position then
// recalculate the recognition table and the parsing tree
// (wihtout concurrency).
func (g *GoCYK) InsertNC(s string, pos int) error {
	c, err := g.Table.Insert(pos)
	if err != nil {
		return err
	}
	g.Sub = append(g.Sub, "")
	copy(g.Sub[pos+1:], g.Sub[pos:])
	g.Sub[pos] = s
	g.completeColumn(s, c, pos)
	if pos < g.Table.Size()-1 {
		g.insertFollowingWithoutC(pos)
	}
	return nil
}

// Remove removes the substring at the given position then
// recalculate the recognition table and the parsing tree.
func (g *GoCYK) Remove(pos int) error {
	if err := g.Table.Remove(pos); err != nil {
		return err
	}
	g.Sub = append(g.Sub[:pos], g.Sub[pos+1:]...)
	g.removeFollowing(pos)
	return nil
}

// IsValid returns true if the parsing string follows the grammar.
func (g *GoCYK) IsValid() bool {
	return g.Table.IsValid()
}
