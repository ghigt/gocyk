package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ghigt/gocyk"
	"github.com/ghigt/gocyk/rtable"
	"github.com/ghigt/gotd/term"
)

var (
	verbose = flag.Bool("v", false, "Print recognition table")
	clear   = flag.Bool("c", false, "Clear print")
	sleep   = flag.Duration("s", 500*time.Millisecond,
		"SleepTime before print")
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

	// init term
	if err := term.TGetEnt(); err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for i := 0; scanner.Scan(); i++ {
		cyk.Add(scanner.Text())
		if *verbose {
			if *clear {
				if err := term.SetCap("cl"); err != nil {
					fmt.Println(err)
				}
			}
			out := rtable.PrettyPrint(cyk.Table, cyk.Sub)
			fmt.Println(out)
			time.Sleep(*sleep)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	cyk.BuildTrees()

	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
		os.Exit(1)
	}
}
