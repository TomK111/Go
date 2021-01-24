package main

import "fmt"

func main() {
	const (
		// EST = -5
		// MST = -7
		// PST = -8
		EST = -(5 + iota) //-5(5+0) = -5
		_                 // -(5+1) = -6
		MST               // -(5+2) = -7
		PST               // -(5+2) = -8
	)

	fmt.Println(EST, MST, PST)
}
