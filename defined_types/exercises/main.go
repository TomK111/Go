package main

import (
	"fmt"
	"workspace/defined_types/exercises/weights"
)

func main() {
	type (
		// Gram's underlying type is int
		Gram int
		// Kilogram's underlying type is int
		Kilogram Gram
		// Ton's underlying type is int
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %v, apples: %v, truck: %v\n", salt, apples, truck)

	// salt = Gram(apples)
	// apples = Kilogram(truck)
	// truck = Ton(Gram(int(apples)))
	salt = Gram(weights.Gram(100))
}
