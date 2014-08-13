package main

import (
	"fmt"
	"time"

	"github.com/ghigt/gocyk"
)

func computeBegNotOnline(input string) {

	// Scanning input
	sub := scanning(input)

	// TO ADD AT THE END
	t := time.Now()

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}

	// Insert at the beginning
	cyk.Insert("package", 0)

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeBegNotOnline\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}

func computeMidNotOnline(input string) {

	// Scanning input
	sub := scanning(input)

	// TO ADD AT THE END
	t := time.Now()

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}
	// Insert in the middle
	cyk.Insert("if", 35)

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeMidNotOnline\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}

func computeEndNotOnline(input string) {

	// Scanning input
	sub := scanning(input)

	// TO ADD AT THE END
	t := time.Now()

	// Instantiate the library
	cyk := gocyk.New(&grammarGo)

	// Compute everything
	for _, s := range sub {
		cyk.Add(s)
	}
	// Add at the end
	cyk.Add("}")

	// Build the Tree
	cyk.BuildTrees()

	fmt.Printf("computeEndNotOnline\t:\t%v\t:\t", time.Since(t))

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}
