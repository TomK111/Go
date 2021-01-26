package main

import (
	"fmt"
)

func main() {
	names := [3]string{}
	fmt.Printf("names: %#v\n", names)
	fmt.Printf("%T\n", names)
	fmt.Printf("%q\n", names)
	fmt.Println()

	distances := [5]int{0, 0, 0, 0}
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("%T\n", distances)
	fmt.Printf("%v\n", distances)
	fmt.Println()

	data := [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("%T\n", data)
	fmt.Printf("%v\n", data)
	fmt.Println()

	ratios := [1]float64{0}
	fmt.Printf("ratios: %#v\n", ratios)
	fmt.Printf("%T\n", ratios)
	fmt.Printf("%v\n", ratios)
	fmt.Println()

	alives := [4]bool{false, false, false, false}
	fmt.Printf("alives: %#v\n", alives)
	fmt.Printf("%T\n", alives)
	fmt.Printf("%v\n", alives)
	fmt.Println()

	zero := [0]uint8{}
	fmt.Printf("zero: %#v\n", zero)
	fmt.Printf("%T\n", zero)
	fmt.Printf("%v\n", zero)
	fmt.Println()

}
