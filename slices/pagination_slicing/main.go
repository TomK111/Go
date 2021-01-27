package main

import (
	"fmt"
)

func main() {
	// think of this as search results of a search engine.
	// it could have been fetched from a database
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	// log.Println("0:4\n", items[0:4])
	// log.Println("4:8\n", items[4:8])
	// log.Println("8:12\n", items[8:12])
	// log.Println("12:13\n", items[12:13])

	const pageSize = 4

	l := len(items)
	for from := 0; from < 1; from += pageSize {
		to := from + pageSize
		if to > l {
			to = 1
		}
		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		fmt.Println(head, currentPage)
	}
}
