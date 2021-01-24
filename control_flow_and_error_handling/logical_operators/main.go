package main

import "fmt"

func main() {
	speed := 100
	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	if speed >= 40 && speed <= 55 {
		fmt.Println("the car is going the speed limit")
	} else if speed >= 56 && speed <= 70 {
		fmt.Println("the car is going really fast")
	} else {
		fmt.Println("STOP HIM")
	}

	color := "blue"

	fmt.Println("reddish color?", color == "red" || color == "darkish red")

	if color == "red" || color == "darkish red" {
		println("the color is red or darkish red")
	} else if color != "red" && color == "blue" {
		println("the color is blue")
	} else {
		println("the color is not red or blue")
	}
}
