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
		`"[a-zA-Z]*"`, // Copty Text
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
	"FuDec": {
		grm.NonTerminal{"FuDe", "PClose"},  // ( )
		grm.NonTerminal{"FuPar", "PClose"}, // ( ... )
	},
	"FuPar":  {grm.NonTerminal{"FuDe", "MulFP"}},
	"FuDe":   {grm.NonTerminal{"FuOp", "POpen"}},
	"FuOp":   {grm.NonTerminal{"FuDef", "Alpha"}},
	"FuDef":  {"func"},
	"FOpen":  {`\{`},
	"FClose": {`\}`},

	// Multiple Function Parameters
	"MulFP": {
		grm.NonTerminal{"Alpha", "Type"}, // Copy Param

		grm.NonTerminal{"MuFPC", "Param"},
	},
	"MuFPC": {grm.NonTerminal{"MulFP", "Comma"}},

	// Parameter
	"Param": {grm.NonTerminal{"Alpha", "Type"}},

	// Body
	"BodyFun": {
		grm.NonTerminal{"FuHead", "Inst"}, // for one or + instructions
		grm.NonTerminal{"FuHead", "If"},   // for one if inst
		grm.NonTerminal{"FuHead", "Var"},  // for one var inst
		grm.NonTerminal{"FuHead", "Call"}, // for one var inst
	},

	// Instruction
	"Inst": {
		grm.NonTerminal{"BodyIf", "FClose"}, // Copy If
		grm.NonTerminal{"IfHead", "FClose"}, // Copy If

		grm.NonTerminal{"VarDec", "Type"}, // Copy Var

		grm.NonTerminal{"CaHead", "PClose"}, // Copy Call
		grm.NonTerminal{"CaDec", "PClose"},  // Copy Call

		grm.NonTerminal{"Inst", "Var"},
		grm.NonTerminal{"Inst", "If"},
		grm.NonTerminal{"Inst", "Call"},
	},

	// If
	"If": {
		grm.NonTerminal{"BodyIf", "FClose"},
		grm.NonTerminal{"IfHead", "FClose"}, // eg. `if test == "test" { }`
	},
	"IfHead": {grm.NonTerminal{"IfDec", "FOpen"}},
	"IfDec":  {grm.NonTerminal{"IfOp", "Val"}},
	"IfOp":   {grm.NonTerminal{"IfDV", "Comp"}},
	"IfDV":   {grm.NonTerminal{"IfDef", "Val"}},
	"IfDef":  {"if"},

	// BodyIf
	"BodyIf": {
		grm.NonTerminal{"IfHead", "Inst"},
	},

	// Comparator
	"Comp": {"==", "!=", ">=", "<=", ">", "<"},

	// Call
	"Call": {
		grm.NonTerminal{"CaHead", "PClose"},
		grm.NonTerminal{"CaDec", "PClose"},
	},
	"CaHead": {grm.NonTerminal{"CaDec", "MulVal"}},
	"CaDec":  {grm.NonTerminal{"MuCaD", "POpen"}},

	"MuCaD": {
		`[a-zA-Z]+`, // Alpha
		grm.NonTerminal{"MCDP", "Alpha"},
	},
	"MCDP": {grm.NonTerminal{"MuCaD", "Point"}},

	// Point
	"Point": {"."},

	// Var
	"Var":    {grm.NonTerminal{"VarDec", "Type"}},
	"VarDec": {grm.NonTerminal{"VarDef", "Alpha"}},
	"VarDef": {"var"},

	// MulVal
	"MulVal": {
		`"[a-zA-Z]*"`, // Copy Val
		`[a-zA-Z]+`,   // Copy Val
		`[0-9]+`,      // Copy Val
		grm.NonTerminal{"CaHead", "PClose"}, // Copy Val
		grm.NonTerminal{"CaDec", "PClose"},  // Copy Val
		`[a-zA-Z]+`,                         // Copy Val
		grm.NonTerminal{"MCDP", "Alpha"},    // Copy Val

		grm.NonTerminal{"MuVaC", "Val"},
	},
	"MuVaC": {grm.NonTerminal{"MulVal", "Comma"}},

	// Comma
	"Comma": {","},

	// Value
	"Val": {
		`"[a-zA-Z]*"`, // Text
		`[a-zA-Z]+`,   // Alpha
		`[0-9]+`,      // Number

		grm.NonTerminal{"CaHead", "PClose"}, // Copy Call
		grm.NonTerminal{"CaDec", "PClose"},  // Copy Call

		`[a-zA-Z]+`,                      // Copy MuCaD
		grm.NonTerminal{"MCDP", "Alpha"}, // Copy MuCaD
	},

	// Type
	"Type": {
		"int", "string", "float",
	},
	// Text
	"Text": {`"[a-zA-Z]*"`},

	// Alphabet
	"Alpha": {"[a-zA-Z]+"},

	// Number (simple)
	"Number": {`[0-9]+`},
}
