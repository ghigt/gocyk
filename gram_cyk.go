package main

import (
	gram "github.com/ghigt/gocyk/grammar"
)

var grammar = gram.Grammar{
	"Number": {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		gram.NonTerminal{"Integer", "Digit"},
		gram.NonTerminal{"N1", "Scale"},
		gram.NonTerminal{"Integer", "Fraction"},
	},
	"N1": {
		gram.NonTerminal{"Integer", "Fraction"},
	},
	"Integer": {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		gram.NonTerminal{"Integer", "Digit"},
	},
	"Fraction": {
		gram.NonTerminal{"T1", "Integer"},
	},
	"T1": {
		".",
	},
	"Scale": {
		gram.NonTerminal{"N2", "Integer"},
	},
	"N2": {
		gram.NonTerminal{"T2", "Sign"},
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
