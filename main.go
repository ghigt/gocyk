package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rule struct {
	Left     Token
	Right    Token
	Terminal string
}

type Rules []*Rule

type Token int

const (
	ILLEGAL Token = iota
	Number
	N1
	Integer
	Fraction
	T1
	Scale
	N2
	T2
	Digit
	Sign
)

var grammar = map[Token]Rules{
	Number: {
		{ILLEGAL, ILLEGAL, "0"},
		{ILLEGAL, ILLEGAL, "1"},
		{ILLEGAL, ILLEGAL, "2"},
		{ILLEGAL, ILLEGAL, "3"},
		{ILLEGAL, ILLEGAL, "4"},
		{ILLEGAL, ILLEGAL, "5"},
		{ILLEGAL, ILLEGAL, "6"},
		{ILLEGAL, ILLEGAL, "7"},
		{ILLEGAL, ILLEGAL, "8"},
		{ILLEGAL, ILLEGAL, "9"},
		{Integer, Digit, ""},
		{N1, Scale, ""},
		{Integer, Fraction, ""},
	},
	N1: {
		{Integer, Fraction, ""},
	},
	Integer: {
		{ILLEGAL, ILLEGAL, "0"},
		{ILLEGAL, ILLEGAL, "1"},
		{ILLEGAL, ILLEGAL, "2"},
		{ILLEGAL, ILLEGAL, "3"},
		{ILLEGAL, ILLEGAL, "4"},
		{ILLEGAL, ILLEGAL, "5"},
		{ILLEGAL, ILLEGAL, "6"},
		{ILLEGAL, ILLEGAL, "7"},
		{ILLEGAL, ILLEGAL, "8"},
		{ILLEGAL, ILLEGAL, "9"},
		{Integer, Digit, ""},
	},
	Fraction: {
		{T1, Integer, ""},
	},
	T1: {
		{ILLEGAL, ILLEGAL, "."},
	},
	Scale: {
		{N2, Integer, ""},
	},
	N2: {
		{T2, Sign, ""},
	},
	T2: {
		{ILLEGAL, ILLEGAL, "e"},
	},
	Digit: {
		{ILLEGAL, ILLEGAL, "0"},
		{ILLEGAL, ILLEGAL, "1"},
		{ILLEGAL, ILLEGAL, "2"},
		{ILLEGAL, ILLEGAL, "3"},
		{ILLEGAL, ILLEGAL, "4"},
		{ILLEGAL, ILLEGAL, "5"},
		{ILLEGAL, ILLEGAL, "6"},
		{ILLEGAL, ILLEGAL, "7"},
		{ILLEGAL, ILLEGAL, "8"},
		{ILLEGAL, ILLEGAL, "9"},
	},
	Sign: {
		{ILLEGAL, ILLEGAL, "+"},
		{ILLEGAL, ILLEGAL, "-"},
	},
}

type Items []Token

func (t Token) String() string {
	switch t {
	case ILLEGAL:
		return fmt.Sprintf("ILLEGAL")
	case Number:
		return fmt.Sprintf("Number")
	case N1:
		return fmt.Sprintf("N1")
	case Integer:
		return fmt.Sprintf("Integer")
	case Fraction:
		return fmt.Sprintf("Fraction")
	case T1:
		return fmt.Sprintf("T1")
	case Scale:
		return fmt.Sprintf("Scale")
	case N2:
		return fmt.Sprintf("N2")
	case T2:
		return fmt.Sprintf("T2")
	case Digit:
		return fmt.Sprintf("Digit")
	case Sign:
		return fmt.Sprintf("Sign")
	}
	return ""
}

func (items *Items) String() string {
	var buf bytes.Buffer

	buf.WriteString("{ ")
	for _, i := range *items {
		buf.WriteString(i.String() + " ")
	}
	buf.WriteString("}")

	return buf.String()
}

var treeTab [][]*Items

func check_grammar(tokenI, tokenJ Token, length, index int) {
	items := treeTab[index][length]

	for index, rules := range grammar {
		for _, rule := range rules {
			if rule.Left == tokenI && rule.Right == tokenJ {
				if items == nil {
					items = new(Items)
				}
				*items = append(*items, Token(index))
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
			if rule.Terminal == s {
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

	// print_tree()
	if treeTab[0][len(input)-1] != nil {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}
}
