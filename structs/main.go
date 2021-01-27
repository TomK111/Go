package main

import "fmt"

//STRUCTS
/*
It's like a blue print or a class in OOP languages like C#
It groups related data in a single type
It's fixed at compile time

Person
-------
Name
LastName
Age

type VideoGame struct {
	Title string
	Genre string
	Published bool
}

pacman := VideoGame {
	Title: "Pac-Man",
	Genre: "Arcade Game"
	Published: true,
}

*/

func main() {
	type Person struct {
		name     string
		lastName string
		age      int
	}

	var picasso Person

	var freud Person

	picasso.name = "Pablo"
	picasso.lastName = "Picasso"
	picasso.age = 91

	freud.name = "Sigmund"
	freud.lastName = "Freud"
	freud.age = 83

	fmt.Printf("\nPicasso: %v\n", picasso)
	fmt.Printf("\nFreud: %v\n", freud)

}
