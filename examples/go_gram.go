package main

import (
	grm "github.com/ghigt/gocyk/grammar"
)

/* Example of a tiny pseudo Golang grammar e.g:

package main

import (
	"fmt"
)

func main ( ) {
	var test int
	if test == 0 {
		fmt . Println ( test )
	}
}

*/
var grammar = grm.Grammar{
	"Program": {grm.NonTerminal{"Pr1", "Function"}},
	"Pr1":     {grm.NonTerminal{"Package", "Import"}},
	// Package
	"Package": {grm.NonTerminal{"Pa1", "Alpha"}},
	"Pa1":     {"package"},
	// Import
	"Import": {grm.NonTerminal{"Im1", "Paren"}},
	"Im1":    {"import"},
	// Function
	"Function": {"func"},
	// Parenthesis
	"Paren":  {grm.NonTerminal{"Par1", "PClose"}},
	"Par1":   {grm.NonTerminal{"POpen", "Text"}},
	"POpen":  {"("},
	"PClose": {")"},
	// Text
	"Text": {`^"[a-z]+"$`},
	// Alphabet
	"Alpha": {"^[a-z]+$"},
}
