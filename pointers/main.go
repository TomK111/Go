package main

import "fmt"

func main() {
	//A pointer stroes the memory address of a value
	//You use ampersand (&) for pointers.
	var (
		counter int
		V       int
		P       *int
	)

	counter = 100
	P = &counter
	V = *P

	fmt.Printf("counter: %-13d address: %-13p\n", counter, &counter)
	fmt.Printf("P : %-13p address: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V: %-13d address: %-13p\n", V, &V)

	// if P == nil {
	// 	fmt.Printf("P is nil and its address is %p\n", P)
	// }

	//Acc: 7933395654
	//Rout: 074908594

	// if P == &counter {
	// 	fmt.Printf("P is equal to counters address: %p == %p\n", P, &counter)
}
