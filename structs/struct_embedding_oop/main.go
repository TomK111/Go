package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn string
	}

	moby := book{
		text: text{title: "Moby Dick", words: 206592},
		isbn: "10234",
	}

	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.title, moby.words, moby.isbn)

}
