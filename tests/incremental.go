package main

import (
	"fmt"
	"time"

	"github.com/ghigt/gocyk"
)

func computeBegIncremental(input string) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	// TO ADD AT THE END
	t := time.Now()

	// Insert at the beginning
	cyk.Insert("package", 0)

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeBegIncremental\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}

func computeMidIncremental(input string) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	// TO ADD AT THE END
	t := time.Now()

	// Insert in the middle
	cyk.Insert("if", 35)

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeMidIncremental\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}

func computeEndIncremental(input string) {

	// Scanning input
	sub := scanning(input)

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	// TO ADD AT THE END
	t := time.Now()

	// Add at the end
	cyk.Add("}")

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeEndIncremental\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}
