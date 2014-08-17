package main

import (
	"time"

	"github.com/ghigt/gocyk"
)

func computeCYK(input []byte) *gocyk.GoCYK {
	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	return cyk
}

func computeTreeNormal(cyk *gocyk.GoCYK) time.Duration {

	t := time.Now()

	// Build the Tree
	cyk.BuildTreesNotC()

	nt := time.Since(t)

	//echo(cyk, nt, "computeTreeNormal")

	return nt
}

func computeTreeConcurrently(cyk *gocyk.GoCYK) time.Duration {

	t := time.Now()

	// Build the Tree
	cyk.BuildTrees()

	nt := time.Since(t)

	//echo(cyk, nt, "computeTreeConcurrently")

	return nt
}
