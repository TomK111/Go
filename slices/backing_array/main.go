package main

import "fmt"

func main() {
	//Slice literal creates a hidden array and returns a slice that refers to that array.
	//The backing array stores the elements not the slice.
	//A Slice doesnt directly store any elements.
	//A slice created with a slice expression share the same array with the original slice.
	//Slices see the changes in the shared backing array.
	//This means if you change the value of the sliced array, the element also changes in backing array.

	grades := [...]float64{40, 10, 20, 50, 60, 70}
	front := grades[:3]
	fmt.Println(front)
	all := grades[:]
	fmt.Printf("all pointer: %p\n", &all)
}
