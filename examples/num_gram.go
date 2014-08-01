package main

import (
	grm "github.com/ghigt/gocyk/grammar"
)

// Example of a numerical grammar (e.g "12.3e+4")
var grammar = grm.Grammar{
	"Number": {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		grm.NonTerminal{"Integer", "Digit"},
		grm.NonTerminal{"N1", "Scale"},
		grm.NonTerminal{"Integer", "Fraction"},
	},
	"N1": {
		grm.NonTerminal{"Integer", "Fraction"},
	},
	"Integer": {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		grm.NonTerminal{"Integer", "Digit"},
	},
	"Fraction": {
		grm.NonTerminal{"T1", "Integer"},
	},
	"T1": {
		".",
	},
	"Scale": {
		grm.NonTerminal{"N2", "Integer"},
	},
	"N2": {
		grm.NonTerminal{"T2", "Sign"},
	},
	"T2": {
		"e",
	},
	"Digit": {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	},
	"Sign": {
		"+",
		"-",
	},
}
