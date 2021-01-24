package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// nil value means the value is not initialized yet
// nil == null(javascript)
// nil == null(c# or java)
// nil gives zero value for pointer-based types such as : pointers, slices, maps, interfaces and channels.

/*
err := do()
if err != nil {
	error
}

//success...continue

*/

const usage = `
Feet to Meters
--------------
This program converts feet into meters

Usage:

feet [feetsToConvert]
`

func main() {
	// n, err := strconv.Atoi(os.Args[1])
	// fmt.Println("Converted number :", n)
	// fmt.Println("Returned error value :", err)

	// age := os.Args[1]
	// n, err := strconv.Atoi(age)
	// if err != nil { //not nill == error
	// 	fmt.Println("ERROR: ", err)
	// 	return
	// }
	// fmt.Printf("Success: Converted %q to %d. \n", age, n)

	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Printf("error: %q is not a number\n", arg)
		return
	}

	meters := feet * 0.3048
	fmt.Printf("%g feet is %g in meters", feet, meters)

}
