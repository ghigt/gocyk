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
var grammarGo = grm.Grammar{
	"Program": {grm.NonTerminal{"Head", "Func"}},
	"Head":    {grm.NonTerminal{"Package", "Import"}},

	// Package
	"Package": {grm.NonTerminal{"PaDef", "Alpha"}},
	"PaDef":   {"package"},

	// Import
	"Import": {grm.NonTerminal{"ImDef", "Paren"}},
	"ImDef":  {"import"},

	// Parenthesis
	"Paren": {grm.NonTerminal{"Par1", "PClose"}},
	"Par1":  {grm.NonTerminal{"POpen", "MulTex"}},
	"MulTex": {
		`"[a-z]+"`,
		grm.NonTerminal{"MulTex", "Text"},
	},
	"POpen":  {`\(`},
	"PClose": {`\)`},

	// Function
	"Func":   {grm.NonTerminal{"BodyDef", "CloseF"}},
	"FuHead": {grm.NonTerminal{"FuDec", "OpenF"}},
	"FuDec":  {grm.NonTerminal{"FuDe", "PClose"}},
	"FuDe":   {grm.NonTerminal{"FuOp", "POpen"}},
	"FuOp":   {grm.NonTerminal{"FuDef", "Alpha"}},
	"FuDef":  {"func"},
	"OpenF":  {`\{`},
	"CloseF": {`\}`},

	// Body
	"BodyDef": {grm.NonTerminal{"FuHead", "Var"}},

	// Var
	"Var":    {grm.NonTerminal{"VarDec", "Type"}},
	"VarDec": {grm.NonTerminal{"VarDef", "Alpha"}},
	"VarDef": {"var"},

	// Type
	"Type": {
		"int", "string", "float",
	},
	// Text
	"Text": {`"[a-z]+"`},

	// Alphabet
	"Alpha": {"[a-z]+"},
}
