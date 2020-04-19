package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The program requires an integer value to run")
		return
	}

	age := os.Args[1]

	n, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}
