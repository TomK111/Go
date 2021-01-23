package main

import "fmt"

func main() {
	scary, speedOfCar := true, 50
	fmt.Println(scary, speedOfCar)

	name, lastName := "Nikola", "Tesla"
	fmt.Println(name, lastName)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	degree, ratio, heat := 10.55, 30.5, 20.0
	fmt.Println(degree, ratio, heat)
}
