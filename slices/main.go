package main

import "fmt"

//Arrays can not grow or shrink. It is a fixed collection type.
//Slices can grow and shrink in runtime. It is a dynamic collection type.

/*
You can not add or remove elements to/from an array in runtime because its length is part of its
type.

It's length belongs to compile time.

var nums []int = slice
var nums [5]int = array

A slice zero value is nil. That means the value has not been initialized.


*/

func main() {
	var books [5]string
	books[0] = "Dracula"
	books[1] = "1984"
	books[2] = "Island"

	games := []string{"Kokemon", "June"}
	newGames := []string{"New Game", "Harry Potter", "Fight Night"}

	newGames = games

	games = nil

	for i, game := range games {
		if game != newGames[i] {
			fmt.Println("not")
			break
		}
	}

	if games == nil {
		fmt.Println("not")
	}
	fmt.Println("gamesOne and newGames are equal")

	fmt.Printf("books: %#v\n", books)
	fmt.Printf("games: %#v\n", games)

	fmt.Printf("games: %T\n", books)
	fmt.Printf("games length: %d\n", len(games))
	fmt.Printf("nil? : %t\n", games == nil)
}
