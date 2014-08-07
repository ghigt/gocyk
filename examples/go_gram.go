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
	"Program": {
		grm.NonTerminal{"Head", "Func"},
		grm.NonTerminal{"Package", "Func"},
	},
	"Head": {grm.NonTerminal{"Package", "Import"}},

	// Package
	"Package": {grm.NonTerminal{"PaDef", "Alpha"}},
	"PaDef":   {"package"},

	// Import
	"Import": {grm.NonTerminal{"ImDef", "Paren"}},
	"ImDef":  {"import"},

	// Parenthesis
	"Paren": {
		grm.NonTerminal{"Par1", "PClose"},
		grm.NonTerminal{"POpen", "PClose"},
	},
	"Par1": {grm.NonTerminal{"POpen", "MulTex"}},
	"MulTex": {
		`"[a-z]+"`, // Text
		grm.NonTerminal{"MulTex", "Text"},
	},
	"POpen":  {`\(`},
	"PClose": {`\)`},

	// Function
	"Func": {
		grm.NonTerminal{"BodyFun", "FClose"},
		grm.NonTerminal{"FuHead", "FClose"},
	},
	"BodyFun": {grm.NonTerminal{"FuHead", "Inst"}},
	"FuHead":  {grm.NonTerminal{"FuDec", "FOpen"}},
	"FuDec":   {grm.NonTerminal{"FuDe", "PClose"}},
	"FuDe":    {grm.NonTerminal{"FuOp", "POpen"}},
	"FuOp":    {grm.NonTerminal{"FuDef", "Alpha"}},
	"FuDef":   {"func"},
	"FOpen":   {`\{`},
	"FClose":  {`\}`},

	// Instruction
	"Inst": {
		grm.NonTerminal{"VarDec", "Type"}, // Var
		grm.NonTerminal{"Inst", "Var"},
		grm.NonTerminal{"IfDec", "FClose"}, // If
		grm.NonTerminal{"Inst", "If"},
	},

	// If
	"If":    {grm.NonTerminal{"IfDec", "FClose"}},
	"IfDec": {grm.NonTerminal{"IfDe", "FOpen"}},
	"IfDe":  {grm.NonTerminal{"IfOp", "Val"}},
	"IfOp":  {grm.NonTerminal{"IfDV", "Comp"}},
	"IfDV":  {grm.NonTerminal{"IfDef", "Alpha"}},
	"IfDef": {"if"},

	// Comparator
	"Comp": {"==", "!=", ">=", "<="},

	// Var
	"Var":    {grm.NonTerminal{"VarDec", "Type"}},
	"VarDec": {grm.NonTerminal{"VarDef", "Alpha"}},
	"VarDef": {"var"},

	// Value
	"Val": {
		`"[a-z]+"`, // Text
		"Number",   // Not Implemented yet
	},

	// Type
	"Type": {
		"int", "string", "float",
	},
	// Text
	"Text": {`"[a-z]+"`},

	// Alphabet
	"Alpha": {"[a-z]+"},
}
