package main

import (
	"fmt"
)

func main() {
	// You can not have a constant with no value.
	//Go can't detect runtime errors at compile-time.
	//Its easier to catch errors with constants

	//RULES
	//1. You can not change constant values (immutable).
	//2. You can not initialize a constant to a run time value.
	//3. vars(variables) belong to run time.

	//4. You can initialize a constant using built-in function len.
	// argument to len is a constant.
	const meters int = 100
	cm := 100
	m := cm / meters
	fmt.Printf("%dcm = %dm\n", cm, m)

	const total int = 5
	const numberOf int = 0

	// fmt.Printf(total / numberOf)
}
