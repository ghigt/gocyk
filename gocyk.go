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
