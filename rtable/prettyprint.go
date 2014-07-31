package rtable

import (
	"fmt"
	"strings"
)

// max returns the meximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// getStringTab returns a table of all the tokens of the recognition
// table converted into strings.
func getStringTab(rt *RTable) (*[][][]string, int) {
	s := make([][][]string, rt.Size())
	m := 1

	for i := 0; i < len(s); i++ {
		col := rt.GetColumn(i)
		s[i] = make([][]string, col.Size())
		for ii := 0; ii < len(s[i]); ii++ {
			item := col.GetItem(ii)
			s[i][ii] = make([]string, item.Size())
			for iii := 0; iii < len(s[i][ii]); iii++ {
				s[i][ii][iii] = string((*item)[iii])
				m = max(len(s[i][ii][iii]), m)
			}
		}
	}
	return &s, m
}

// printToken prints a string followed by spaces corresponding to
// the `max` parameter.
func printToken(s string, max int) {
	fmt.Printf(s)
	if l := len(s); l < max {
		fmt.Printf(strings.Repeat(" ", max-l))
	}
}

// PrettyPrint prints a given recognition table.
func PrettyPrint(rt *RTable) string {
	s, m := getStringTab(rt)
	length := rt.Size()

	fmt.Printf("+")
	fmt.Printf(strings.Repeat("-", (m+1)*length-1))
	fmt.Println("+")
	for row := 0; row < length; row++ {
		l := 1
		for tok := 0; tok < l; tok++ {
			fmt.Printf(strings.Repeat(" ", (m+1)*row))
			fmt.Printf("|")
			for col := row; col < length; col++ {
				if len((*s)[col][row]) == 0 {
					printToken("", m)
				} else if tok == 0 {
					l = max(len((*s)[col][row]), l)
					printToken((*s)[col][row][0], m)
				} else if tok < len((*s)[col][row]) {
					printToken((*s)[col][row][tok], m)
				} else {
					printToken("", m)
				}
				fmt.Printf("|")
			}
			fmt.Println()
		}
		fmt.Printf(strings.Repeat(" ", (m+1)*row))
		fmt.Printf("+")
		fmt.Printf(strings.Repeat("-", (m+1)*(length-row)-1))
		if row+1 == length {
			fmt.Println("+")
		} else {
			fmt.Println("|")
		}
	}

	return ""
}
