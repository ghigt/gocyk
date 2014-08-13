package main

import (
	"fmt"
	"time"

	"github.com/ghigt/gocyk"
)

func computeBegOnline(input string) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	t := time.Now()

	// Insert at the beginning
	cyk.Add("package")

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeBegOnline\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}

func computeMidOnline(input string) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	var t time.Time

	// Compute everything
	for i, s := range sub {
		if i == 35 {
			t = time.Now()

			// Insert in the middle
			cyk.Add("if")
		}
		cyk.Add(s)
	}

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeMidOnline\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}

func computeEndOnline(input string) {

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
	cyk.Add("}")

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeEndOnline\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}
