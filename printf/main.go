package main

import "fmt"

func main() {
	var brand string = "Google"

	fmt.Printf("%s\n", brand)  //Prints string without quotes
	fmt.Printf("%q\n", brand)  //Prints string with quotes
	fmt.Printf("%T\n", brand)  //Prints data type of value
	fmt.Printf("%#v\n", brand) //Prints value of variable
	fmt.Printf("%v\n", brand)  //Prints value of variable without quotes(string)

	fmt.Println("hi\nhi") // \n is an escape sequence that prints a new line.

	var speed int
	var head float64
	var off bool
	var brands string

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", head)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brands)

	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d\n", distance)
	fmt.Printf("Orbital: %f\n", orbital)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)
	fmt.Printf(
		"%v is %v kms away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance,
	)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Orbital Period: %.3f days\n", orbital)
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
	fmt.Printf("Orbital Period: %.5f days\n", orbital)
	fmt.Printf("Orbital Period: %.6f days\n", orbital)
	fmt.Printf("Orbital Period: %.7f days\n", orbital)

}
