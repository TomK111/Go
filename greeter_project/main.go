package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	name = os.Args[0]
	fmt.Println("Hello great", name, "!")
	name, age := "Gandalf", 2019

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall not pass!")
}
