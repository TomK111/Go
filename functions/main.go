package main

import "fmt"

var N int = 1

func main() {
	showN()
}

func showN() {
	if N == 0 {
		return
	}
	fmt.Printf("showN : N is %d\n", N)
}

func incrN() {
	N++
}
