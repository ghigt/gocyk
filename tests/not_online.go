package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ghigt/gocyk"
)

func main() {
	var input string
	cyk := gocyk.New(&grammarGo)

	flag.Parse()

	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	} else {
		fmt.Println("Please, precise an input")
		return
	}

	// TODO: still not a good NOT online or incremental

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for i := 0; scanner.Scan(); i++ {
		cyk.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))

	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
		os.Exit(1)
	}
}
