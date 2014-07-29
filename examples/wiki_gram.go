package main

import (
	grm "github.com/ghigt/gocyk/grammar"
)

// Example grammar from Wikipedia:
// http://en.wikipedia.org/wiki/CYK_algorithm
var grammarWiki = grm.Grammar{
	"S": {
		grm.NonTerminal{"NP", "VP"},
	},
	"VP": {
		"eats",
		grm.NonTerminal{"VP", "PP"},
		grm.NonTerminal{"V", "NP"},
	},
	"PP": {
		grm.NonTerminal{"P", "NP"},
	},
	"NP": {
		"she",
		grm.NonTerminal{"Det", "N"},
	},
	"V": {
		"eats",
	},
	"P": {
		"with",
	},
	"N": {
		"fish", "fork",
	},
	"Det": {
		"a",
	},
}
