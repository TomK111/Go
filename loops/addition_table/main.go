package main

import (
	"fmt"
)

const max = 12

func main() {
	fmt.Println(`                                                 ADDITION TABLE`)
	fmt.Printf("%8s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%8d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%8d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%8d", i+j)
		}
		fmt.Println()
	}

	fmt.Println(`                                                 SUBTRACTION TABLE`)
	fmt.Printf("%8s", "X")
	for k := 0; k < max; k++ {
		fmt.Printf("%8d", k)
	}
	fmt.Println()

	for k := 0; k <= max; k++ {
		fmt.Printf("%8d", k)

		for z := 0; z <= max; z++ {
			fmt.Printf("%8d", k-z)
		}

		fmt.Println()
	}

}
