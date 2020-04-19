package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide an argument")
		return
	}

	max, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("Ensure argument is a valid integer")
		return
	}

	var uniques []int

loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)
	}

	fmt.Println("\n\nuniques:", uniques)

	sort.Ints(uniques)
	fmt.Println("sorted:", uniques)

}
