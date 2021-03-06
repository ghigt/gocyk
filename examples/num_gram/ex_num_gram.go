package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ghigt/gocyk"
	"github.com/ghigt/gocyk/rtable"
	"github.com/ghigt/gotd/term"
)

var (
	verbose = flag.Bool("v", false, "Print recognition table")
	sleep   = flag.Duration("s", 500*time.Millisecond,
		"SleepTime before print")
)

func main() {
	var input string
	cyk := gocyk.New(&grammar)

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

	for scanner.Scan() {
		cyk.Add(scanner.Text())
		if *verbose {
			// clear screen
			if err := term.SetCap("cl"); err != nil {
				fmt.Println(err)
			}
			out := rtable.PrettyPrint(cyk.Table, cyk.Sub)
			fmt.Println(out)
			time.Sleep(*sleep)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	// --INSERT TEST--
	//if err := cyk.Insert(".", 1); err != nil {
	//	log.Fatal(err)
	//}
	//if err := term.SetCap("cl"); err != nil {
	//	fmt.Println(err)
	//}
	//out := rtable.PrettyPrint(cyk.Table, cyk.Sub)
	//fmt.Println(out)
	//if err := cyk.Remove(0); err != nil {
	//	log.Fatal(err)
	//}
	//time.Sleep(*sleep)
	//if err := term.SetCap("cl"); err != nil {
	//	fmt.Println(err)
	//}
	//out = rtable.PrettyPrint(cyk.Table, cyk.Sub)
	//fmt.Println(out)
	// --TEST--

	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		log.Fatal("It fails :(")
	}
}
