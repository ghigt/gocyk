package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ghigt/gocyk/table"
	"github.com/ghigt/gotd/term"
)

var (
	verbose = flag.Bool("v", false, "Print recognition table")
)

func main() {
	var input string
	rtable := table.New(&grammar)

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
	scanner.Split(bufio.ScanBytes)

	for i := 0; scanner.Scan(); i++ {
		rtable.Add(scanner.Text())
		if *verbose {
			// clear screen
			if err := term.SetCap("cl"); err != nil {
				fmt.Println(err)
			}
			table.PrettyPrint(rtable)
			time.Sleep(500 * time.Millisecond)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	// --TEST--
	//if err := rtable.Insert(".", 1); err != nil {
	//	log.Fatal(err)
	//}
	//if err := term.SetCap("cl"); err != nil {
	//	fmt.Println(err)
	//}
	//table.PrettyPrint(rtable)
	//if err := rtable.Remove(1); err != nil {
	//	log.Fatal(err)
	//}
	//if err := term.SetCap("cl"); err != nil {
	//	fmt.Println(err)
	//}
	//table.PrettyPrint(rtable)
	if rtable.ValidFor(1, 4) {
		fmt.Println("Valid from 1 to 4 :)")
	}
	// --TEST--

	if rtable.Valid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}
}