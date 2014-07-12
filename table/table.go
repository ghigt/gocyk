package table

import (
	"bytes"
	"errors"
	"strconv"

	gram "github.com/ghigt/gocyk/grammar"
)

// RTable is the type to build a recognition table
type RTable [][]*gram.Items

func (rt *RTable) Get(index, length int) (*gram.Items, error) {
	if rt == nil {
		return nil, errors.New("Error: null pointer")
	}
	items := (*rt)[index][length]
	return items, nil
}

func (rt *RTable) Add(s string) error {
	if rt == nil {
		return errors.New("Error: null pointer")
	}
	// -- TEST --
	lengthT := make([]*gram.Items, len(*rt)+1)
	lengthT[len(*rt)] = lengthT[len(*rt)].Add(gram.Token(s))
	*rt = append(*rt, lengthT)
	// -- TEST --

	return nil
}

func (rt *RTable) Insert(index int) error {
	if rt == nil {
		return errors.New("Error: null pointer")
	}
	return nil
}

func (rt *RTable) Remove(index int) error {
	if rt == nil {
		return errors.New("Error: null pointer")
	}
	return nil
}

func (rt *RTable) Valid() bool {
	if rt == nil {
		return false
	}
	if len(*rt) == 0 || (*rt)[len(*rt)-1][0] != nil {
		return true
	}
	return false
}

func (rt *RTable) ValidFor(beg, end int) bool {
	if rt == nil {
		return false
	}
	return false
}

func (rt *RTable) String() string {
	var buf bytes.Buffer

	if rt == nil {
		return "nil"
	}
	buf.WriteString("[length][index]\n")
	for i, index := range *rt {
		for l := len(index) - 1; l >= 0; l-- {
			buf.WriteString("[" + strconv.Itoa(l) + "]" +
				"[" + strconv.Itoa(i) + "]")
			if index[l] == nil {
				buf.WriteString("<nil> ")
			} else {
				buf.WriteString(index[l].String() + " ")
			}
		}
		buf.WriteString("\n")
	}

	return buf.String()
}
