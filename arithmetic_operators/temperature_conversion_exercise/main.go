package main

import "fmt"

func main() {
	celsius := 10.2
	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%v Celsius is %v Fahrenheit", celsius, fahrenheit)
}
