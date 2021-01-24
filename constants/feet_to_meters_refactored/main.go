package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const (
		feetInMeters float64 = 0.3048
		feetInYards  float64 = feetInMeters / 0.9144
	)

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := math.Round(feet * feetInMeters)
	yards := math.Round(feet * feetInYards)

	fmt.Printf("%v feet is %v meters\n", feet, meters)
	fmt.Printf("%v feet is %v yards\n", feet, yards)
}
