package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanning(input string) []string {
	sub := []string{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for i := 0; scanner.Scan(); i++ {
		sub = append(sub, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return sub
}
