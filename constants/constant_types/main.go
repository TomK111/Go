package main

import (
	"fmt"
)

//You can use expressions when initializing constants.

func main() {
	const min = 1 + 1
	const pi = 3.14 * min
	const version = "2.0.1"
	const debug = !true

	fmt.Println(min, pi, version, debug)
}
