package main

import "fmt"

func main() {
	fmt.Println(
		-200, -100, 0, 50, 100,
	) //Prints integers

	fmt.Println(
		-50.5, -20.5, -0., 1., 100.56,
	) //Prints floats

	fmt.Println(
		true, false,
	) //Prints boolean values

	fmt.Println(0x0, 0x1, 0x2, 0x3) //Prints hexadecimal values

	fmt.Println(
		"", //empty string
		"hi",
	) //Prints string

}
