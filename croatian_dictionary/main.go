package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide an English word to translate into Croatian.")
		return
	}

	query := args[0]

	dict := map[string]string{}
	dict["Up"] = "Gore"
	dict["Down"] = "Dolje"
	dict["Morning"] = "Dobro Jutro"
	dict["Night"] = "Noc"
	dict["Hello"] = "Zdravo / Bok / Cao / Halo / Dobar Dan"
	dict["How"] = "How"

	value, ok := dict[query]
	if !ok {
		fmt.Printf("%q not found\n", query)
		return
	}

	fmt.Printf("%q is %#v in Croatian\n", query, value)

	fmt.Printf("Zero Value: %#v\n", len(dict))
}
