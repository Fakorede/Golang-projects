package main

import (
	"fmt"
	"os"
	"strconv"
)

const errMsg = "ERROR: Kindly specify the range of multiplication table to be generated."

func main() {

	// ensures an argument is passed
	if len(os.Args) != 2 {
		fmt.Println(errMsg)
		return
	}

	// converts to int
	n, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("%5s", "X")

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}

	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}
}
