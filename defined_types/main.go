package main

import (
	"fmt"
)

type gram float64
type ounce float64

func main() {
	var g gram = 1000
	var o ounce
	o = ounce(g) * 0.035274
	fmt.Printf("%g grams is %.2f ounces\n", g, o)

	// type Duration int64
	// //Create a new type called duration based on int64
	// //int64 is the underlying or real type
	// var ms int64 = 1000
	// var ns Duration
	// ns = Duration(ms)
	// ms = int64(ns)

	// h, _ := time.ParseDuration("4h30m")
	// fmt.Printf("Type of H: %T\n", h)
	// fmt.Printf("Type pf h's underlying type: %T\n", int64(h))

	// fmt.Println(h.Hours(), "hours")
	// fmt.Println(int64(h))

}
