package main

import (
	"fmt"
	"strconv"
)

func main() {
	name, lastName := "carl", "sagan"

	fmt.Println(name + " " + lastName)

	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)

	fmt.Println(eq)
}
