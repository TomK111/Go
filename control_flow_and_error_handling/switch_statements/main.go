package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo", "Nagasaki":
		fmt.Println("Japan")
	default:
		fmt.Println("I am sorry. I can not find the location of that city. Please try again")
	}

	switch i := 10; true {
	case i > 0:
		fmt.Println("Positive")
	case i < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Neither negative nor positive. It is 0.")
	}

	//Behind the scenes, Go converts switch to if statements.

	//RULES
	//1. Case conditions must be unique.
	//2. Types should be comparable.
	//3. Breaks are not necessary.
}
