package rtable

import (
	grm "github.com/ghigt/gocyk/grammar"
)

// TableST is the structure containing the Recognition table and
// the associated grammar.
type TableST struct {
	*RTable
	Grammar *grm.Grammar
}

// New instanciate the structure and the recognition table with
// the given grammar.
func New(g *grm.Grammar) *TableST {
	return &TableST{RTable: &RTable{}, Grammar: g}
}
