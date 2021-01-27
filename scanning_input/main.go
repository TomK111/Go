package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// in := bufio.NewScanner(os.Stdin)

	// var lines int

	// for in.Scan() {
	// 	lines++

	// }
	// fmt.Printf("There are %d lines(s)\n", lines)

	// if err := in.Err(); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter a word.")
		return
	}
	query := args[0]

	rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	if _, ok := words[query]; ok {
		fmt.Printf("The input contains %q\n:", query)
	}
	fmt.Printf("The input does not contain %q\n:", query)
}
