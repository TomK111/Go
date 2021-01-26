package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythings Hip",
		"Kafka's Revenge Second Edition",
	}

	// books[0] = "Kafka's Revenge"
	// books[1] = "Stay Golden"
	// books[2] = "Everythings hip"
	// books[3] = "Kafka's Revenger 2nd Edition"

	fmt.Printf("books : %#v\n", books)
	fmt.Println(books)
}
