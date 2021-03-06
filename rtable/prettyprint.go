package rtable

import (
	"bytes"
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
func printToken(s string, max int, buf *bytes.Buffer) {
	buf.WriteString(s)
	if l := len(s); l < max {
		buf.WriteString(strings.Repeat(" ", max-l))
	}
}

func printSub(sub []string, max int, buf *bytes.Buffer) {
	buf.WriteString("+" + strings.Repeat("-",
		len(sub)*max+len(sub)-1) + "+\n")
	buf.WriteString("|")
	for _, s := range sub {
		printToken(s, max, buf)
		buf.WriteString("|")
	}
	buf.WriteString("\n+" + strings.Repeat("-",
		len(sub)*max+len(sub)-1) + "+\n")
}

// PrettyPrint prints a given recognition table.
func PrettyPrint(rt *RTable, sub []string) string {
	var buf bytes.Buffer

	if rt.Size() == 0 {
		return ""
	}
	s, m := getStringTab(rt)
	length := rt.Size()

	buf.WriteString("+" + strings.Repeat("-",
		(m+1)*length-1) + "+\n")
	for row := 0; row < length; row++ {
		l := 1
		for tok := 0; tok < l; tok++ {
			buf.WriteString(strings.Repeat(" ", (m+1)*row) + "|")
			for col := row; col < length; col++ {
				if len((*s)[col][row]) == 0 {
					printToken("", m, &buf)
				} else if tok == 0 {
					l = max(len((*s)[col][row]), l)
					printToken((*s)[col][row][0], m, &buf)
				} else if tok < len((*s)[col][row]) {
					printToken((*s)[col][row][tok], m, &buf)
				} else {
					printToken("", m, &buf)
				}
				buf.WriteString("|")
			}
			buf.WriteString("\n")
		}
		buf.WriteString(strings.Repeat(" ",
			(m+1)*row) + "+" + strings.Repeat("-",
			(m+1)*(length-row)-1))
		if row+1 == length {
			buf.WriteString("+\n")
		} else {
			buf.WriteString("|\n")
		}
	}

	printSub(sub, m, &buf)

	return buf.String()
}
