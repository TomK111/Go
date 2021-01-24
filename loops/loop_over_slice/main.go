package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)

	for j := 0; j < len(words); j++ {
		fmt.Printf("%q\n", words[j])
	}

	newSentence := strings.Fields(
		"This is a brand new sentence",
	)

	for k := 0; k < len(newSentence); k++ {
		fmt.Printf("%q\n", newSentence[k])
	}

	sentenceTwo := strings.Fields(
		"This is sentence number three actually",
	)

	for index, value := range sentenceTwo {
		fmt.Println(index, value)
	}
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%q\n", os.Args[i])
	// }
}
