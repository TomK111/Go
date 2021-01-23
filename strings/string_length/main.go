package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "carl"

	fmt.Println(len(name))
	//len returns amount of bytes not actual amount of characters.
	//This may be inaccurate when dealing with foreign names(foreign letters)
	fmt.Println(utf8.RuneCountInString(name))
}
