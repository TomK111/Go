package main

import (
	"fmt"
	"time"
)

func main() {
	h := time.Now().Hour()
	fmt.Println(h)

	switch {
	case h >= 18:
		fmt.Println("Good Evening")
	case h >= 12:
		fmt.Println("Good Afternoon")
	case h >= 6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Goodnight")
	}

}
