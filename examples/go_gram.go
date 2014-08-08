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
		grm.NonTerminal{"PaDef", "Alpha"},  // Package alone
		grm.NonTerminal{"Head", "MulF"},    // P+I+MulF
		grm.NonTerminal{"Package", "MulF"}, // Only Package and MulF
	},
	"Head": {grm.NonTerminal{"Package", "Import"}},

	"MulF": {
		grm.NonTerminal{"BodyFun", "FClose"}, // Copy Func
		grm.NonTerminal{"FuHead", "FClose"},  // Copy Func

		grm.NonTerminal{"MulF", "Func"},
	},

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
		`"[a-z]+"`, // Copty Text
		grm.NonTerminal{"MulTex", "Text"},
	},
	"POpen":  {`\(`},
	"PClose": {`\)`},

	// Function
	"Func": {
		grm.NonTerminal{"BodyFun", "FClose"}, // Func not empty
		grm.NonTerminal{"FuHead", "FClose"},  // eg. `func main() { }`
	},
	"FuHead": {grm.NonTerminal{"FuDec", "FOpen"}},
	"FuDec":  {grm.NonTerminal{"FuDe", "PClose"}},
	"FuDe":   {grm.NonTerminal{"FuOp", "POpen"}},
	"FuOp":   {grm.NonTerminal{"FuDef", "Alpha"}},
	"FuDef":  {"func"},
	"FOpen":  {`\{`},
	"FClose": {`\}`},

	// Body
	"BodyFun": {
		grm.NonTerminal{"FuHead", "Inst"}, // for one or + instructions
		grm.NonTerminal{"FuHead", "If"},   // for one if inst
		grm.NonTerminal{"FuHead", "Var"},  // for one var inst
	},

	// Instruction
	"Inst": {
		grm.NonTerminal{"BodyIf", "FClose"}, // Copy If
		grm.NonTerminal{"IfHead", "FClose"}, // Copy If

		grm.NonTerminal{"VarDec", "Type"}, // Var
		grm.NonTerminal{"Inst", "Var"},
		grm.NonTerminal{"Inst", "If"},
	},

	// If
	"If": {
		grm.NonTerminal{"BodyIf", "FClose"},
		grm.NonTerminal{"IfHead", "FClose"}, // eg. `if test == "test" { }`
	},
	"IfHead": {grm.NonTerminal{"IfDec", "FOpen"}},
	"IfDec":  {grm.NonTerminal{"IfOp", "Val"}},
	"IfOp":   {grm.NonTerminal{"IfDV", "Comp"}},
	"IfDV":   {grm.NonTerminal{"IfDef", "Alpha"}},
	"IfDef":  {"if"},

	// BodyIf
	"BodyIf": {
		grm.NonTerminal{"IfHead", "Inst"}, //
	},

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
