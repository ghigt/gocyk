package main

import (
	"time"

	"github.com/ghigt/gocyk"
)

func computeBegIncremental(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	t := time.Now()

	// Insert at the beginning
	cyk.Insert(file.Sub, file.Pos)

	nt := time.Since(t)

	echo(cyk, nt, "computeBegIncremental")
}

func computeMidIncremental(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	t := time.Now()

	// Insert in the middle
	cyk.Insert(file.Sub, file.Pos)

	nt := time.Since(t)

	echo(cyk, nt, "computeMidIncremental")
}

func computeEndIncremental(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	t := time.Now()

	// Add at the end
	cyk.Add(file.Sub)

	nt := time.Since(t)

	echo(cyk, nt, "computeEndIncremental")
}
