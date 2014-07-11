package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var grammar = Grammar{
	"Number": {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		NonTerminal{"Integer", "Digit"},
		NonTerminal{"N1", "Scale"},
		NonTerminal{"Integer", "Fraction"},
	},
	"N1": {
		NonTerminal{"Integer", "Fraction"},
	},
	"Integer": {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		NonTerminal{"Integer", "Digit"},
	},
	"Fraction": {
		NonTerminal{"T1", "Integer"},
	},
	"T1": {
		".",
	},
	"Scale": {
		NonTerminal{"N2", "Integer"},
	},
	"N2": {
		NonTerminal{"T2", "Sign"},
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

var treeTab [][]*Items

func check_grammar(tokenI, tokenJ Token, length, index int) {
	items := treeTab[index][length]

	for i, rules := range grammar {
		for _, rule := range rules {
			v, ok := rule.(NonTerminal)
			if ok && v.Left == tokenI && v.Right == tokenJ {
				if items == nil {
					items = new(Items)
				}
				*items = append(*items, Token(i))
			}
		}
	}
	treeTab[index][length] = items
}

func try_combinaisons(length, index int) {
	for l := 0; l < length; l++ { // loop over lengths
		if treeTab[index][l] != nil {
			for _, tokenI := range *(treeTab[index][l]) { // loop vertical rule
				if treeTab[index+l+1][length-l-1] != nil {
					for _, tokenJ := range *(treeTab[index+l+1][length-l-1]) { // loop diagonal rule
						check_grammar(tokenI, tokenJ, length, index)
					}
				}
			}
		}
	}
}

func fill_tree() {
	for l := 1; l < len(treeTab); l++ {
		for i := 0; i < len(treeTab)-l; i++ {
			try_combinaisons(l, i)
		}
	}
}

func first_length(i int, s string) {
	var items Items

	for index, rules := range grammar {
		for _, rule := range rules {
			v, ok := rule.(Terminal)
			if ok && v == Terminal(s) {
				items = append(items, Token(index))
			}
		}
	}
	treeTab[i][0] = &items
}

func init_tree(input string) {
	treeTab = make([][]*Items, len(input))
	for i := 0; i < len(input); i++ {
		treeTab[i] = make([]*Items, len(input)-i)
	}
}

func print_tree() {
	for _, index := range treeTab {
		for _, length := range index {
			fmt.Printf("%v ", length)
		}
		fmt.Println()
	}
}

func main() {
	var input string

	if len(os.Args) >= 2 {
		input = os.Args[1]
	} else {
		log.Fatal("Please precise an input")
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanBytes)

	init_tree(input)

	for i := 0; scanner.Scan(); i++ {
		first_length(i, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fill_tree()

	print_tree()
	if treeTab[0][len(input)-1] != nil {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}
}
