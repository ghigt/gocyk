package main

import (
	"time"

	"github.com/ghigt/gocyk"
)

func computeBegNotOnline(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	t := time.Now()

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Insert at the beginning
	cyk.Add(file.Sub)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	nt := time.Since(t)

	echo(cyk, nt, "computeBegNotOnline")
}

func computeMidNotOnline(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	t := time.Now()

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}
	// Insert in the middle
	cyk.Insert(file.Sub, file.Pos)

	nt := time.Since(t)

	echo(cyk, nt, "computeMidNotOnline")
}

func computeEndNotOnline(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	t := time.Now()

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}
	// Add at the end
	cyk.Add(file.Sub)

	nt := time.Since(t)

	echo(cyk, nt, "computeEndNotOnline")
}
