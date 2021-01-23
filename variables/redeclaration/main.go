package main

import "fmt"

func main() {
	var safe bool = false
	safe, speed := true, 50
	speed = 70
	safe = false
	speed = 80
	fmt.Println(safe, speed)

	name := "Nikola"
	name, age := "Marie", 65
	fmt.Println(name, age)

	name = "Albert"
	birth := 1879
	fmt.Println(name, birth)

}
