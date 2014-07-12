package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ghigt/gocyk/table"
)

func main() {
	var input string
	rtable := new(table.RTable)

	if len(os.Args) >= 2 {
		input = os.Args[1]
	} else {
		fmt.Println("Please, precise an input")
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanBytes)

	for i := 0; scanner.Scan(); i++ {
		rtable.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Println(rtable)

	if rtable.Valid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}
}
