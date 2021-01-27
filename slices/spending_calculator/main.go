package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1st day: $200, 100$
	//2nd day : $500
	//3rd day: $50,$25, $75

	// spendings := [][]int{
	// 	{200, 100},
	// 	{500},
	// 	{50, 25, 75},
	// }

	// spendings := make([][]int, 0, 5)

	// spendings = append(spendings, []int{200, 100})
	// spendings = append(spendings, []int{25, 10, 45, 50})
	// spendings = append(spendings, []int{5, 15, 40})
	// spendings = append(spendings, []int{95, 10}, []int{51, 59})

	spendings := fetch()

	for i, daily := range spendings {
		var total int

		for _, spending := range daily {
			total += spending
		}
		fmt.Printf("Day %d: %d\n", i+1, total)
	}

}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`

	lines := strings.Split(content, "\n")

	spendings := make([][]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {

			spending, _ := strconv.Atoi(field)

			spendings[i][j] = spending
		}
	}

	return spendings
}
