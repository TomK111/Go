package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	var uniques []int
loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)
	}
	fmt.Printf("\nRandomized Numbers: %v\n", uniques)
	elapsedForRandom := time.Since(start)

	log.Printf("Randomization took: %s", elapsedForRandom)
	sort.Ints(uniques)

	fmt.Printf("\nSorted: %v\n", uniques)
	elapsedForSort := time.Since(start)

	log.Printf("Sorting took %s", elapsedForSort)
	println()
}
