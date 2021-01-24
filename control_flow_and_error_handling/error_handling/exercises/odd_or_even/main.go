package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		println("Pick a number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a number", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%q is divisible by 8", os.Args[1])
	} else if n%2 == 0 {
		fmt.Printf("%q is divisble by 2", os.Args[1])
	} else {
		fmt.Printf("%q is an odd number.")
	}
}
