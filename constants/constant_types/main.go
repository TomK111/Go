package main

import (
	"fmt"
	"time"
)

//You can use expressions when initializing constants.

func main() {
	const min = 1 + 1
	const pi = 3.14 * min
	const version = "2.0.1"
	const debug = !true

	fmt.Println(min, pi, version, debug)

	// i := 10
	// t := time.Second * time.Duration(i)
	// fmt.Println(t)
	//Prints 10s

	const c = 10
	t := time.Second * time.Duration(c)
	fmt.Println(t)
	//Prints 10s
}
