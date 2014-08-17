package main

import (
	"time"

	"github.com/ghigt/gocyk"
)

func computeBegOnline(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	t := time.Now()

	// Insert at the beginning
	cyk.Add(file.Sub)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	nt := time.Since(t)

	echo(cyk, nt, "computeBegOnline")
}

func computeMidOnline(file File, input []byte) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	var t time.Time

	// Compute everything
	for i, s := range sub {
		if i == file.Pos {
			t = time.Now()

			// Insert in the middle
			cyk.Add(file.Sub)
		}
		cyk.Add(s)
	}

	nt := time.Since(t)

	echo(cyk, nt, "computeMidOnline")
}

func computeEndOnline(file File, input []byte) {

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

	echo(cyk, nt, "computeEndOnline")
}
