package main

import "fmt"

const max = 12

func main() {
	fmt.Printf("%8s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%8d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%8d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%8d", i-j)
		}
		fmt.Println()
	}
}
