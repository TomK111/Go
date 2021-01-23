package main

import (
	"fmt"
)

func main() {
	speed := 100
	force := 2.5
	speed = int(float64(speed) * force)
	fmt.Println(speed)

	var apple int
	var orange int32

	orange = 65
	color := string(orange)
	fmt.Println(color) //This will print the letter A.

	apple = int(orange)
	orange = int32(apple)
	fmt.Println(orange)
}
