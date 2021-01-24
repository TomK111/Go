package main

import (
	"fmt"
	"strconv"
)

func main() {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("There was no error, n is", n)
	}
}

//Simple statements allow you to execute a statement inside another statement.
//Declared variables are only visible inside the if statement and its branches.
