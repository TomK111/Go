package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		radius = 10.0
		area   float64
	)

	// area = math.Pi * radius * radius

	//math.Pow calculates the power of a float number.
	area = math.Pi * math.Pow(radius, 2)
	fmt.Printf("radius: %g -> area : %.2f\n", radius, area)

}
