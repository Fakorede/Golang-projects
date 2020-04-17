package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1st day: $200, $100
	// 2nd day: $500
	// 3rd day: $50, $25, and $75

	// spendings := [][]int{
	// 	{200, 100},   // 1st day
	// 	{500},        // 2nd day
	// 	{50, 25, 75}, // 3rd day
	// }

	// for i, daily := range spendings {
	// 	var total int
	// 	for _, spending := range daily {
	// 		total += spending
	// 	}

	// 	fmt.Printf("Day %d: %d\n", i+1, total)
	// }

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

	// splits content line-by-line
	lines := strings.Split(content, "\n")

	// creates a multi-dimensional slice with length of line(5)
	spendings := make([][]int, len(lines))

	// inserts each line into the multi-dimensional slice
	for i, line := range lines {

		// converts each line into slice of string
		fields := strings.Fields(line)

		// creates a slice for each field
		spendings[i] = make([]int, len(fields))

		// converts each field element to an integer
		for j, field := range fields {
			spending, _ := strconv.Atoi(field)

			spendings[i][j] = spending
		}
	}

	return spendings

}
