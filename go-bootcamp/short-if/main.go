package main

import (
	"fmt"
	"os"
	"strconv"
)

// scope of simple statements
func main() {
	if a := os.Args; len(a) != 2 {
		// only a variable
		fmt.Println("Provide a number")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// a, n and err variables
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		// all the variables in the if stmt
		fmt.Printf("%s * 2 %d\n", a[1], n*2)
	}

	//fmt.Printf("%s * 2 %d\n", a[1], n*2) //variables outside the block-scope, won't work.
}
