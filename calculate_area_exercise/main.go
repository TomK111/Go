package main

import "fmt"

func main() {
	width, height := 5.0, 12.0

	area := width * height

	fmt.Println(area)

	area -= 10 // decrease by 10

	area += 10 //increase by 10

	area *= 2 // multiply by 2

	area /= 2 // divide by 2

	fmt.Printf("%g\n", area)
}
