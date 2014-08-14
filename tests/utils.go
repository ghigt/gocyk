package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/ghigt/gocyk"
)

func getFile(name string, prefix string) []byte {
	content, err := ioutil.ReadFile(name + "_" + prefix + ".go.g")

	if err != nil {
		log.Print(err)
	}
	return content
}

func echo(cyk *gocyk.GoCYK, t time.Duration, s string) {
	fmt.Printf("%s\t:\t%v\t:\t", s, t)

	// Check if it works
	if cyk.IsValid() {
		fmt.Println("It works :)")
	} else {
		fmt.Println("It fails :(")
	}

	//fmt.Println(rtable.PrettyPrint(cyk.Table, cyk.Sub))
}

func scanning(content []byte) []string {
	sub := []string{}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	scanner.Split(bufio.ScanWords)

	for i := 0; scanner.Scan(); i++ {
		sub = append(sub, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return sub
}
