package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	celsius, _ := strconv.ParseFloat(arg, 64)

	fahrenheight := (9*celsius + 160) / 5

	fmt.Printf("%g Celsius is %g Fahrenheight", celsius, fahrenheight)
}
