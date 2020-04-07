package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	errMsg1 = "ERROR: Kindly specify the range of multiplication table to be generated."
	errMsg2 = "ERROR: kindly ensure the passed argument is of type int."
)

func main() {

	// ensures an argument is passed
	if len(os.Args) != 2 {
		fmt.Println(errMsg1)
		return
	}

	// converts argument to int
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(errMsg2)
		return
	}

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
