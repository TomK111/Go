package main

import (
	"fmt"
)

func main() {
	//iota is a built-on constant generator which generates every increasing numbers.
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}
