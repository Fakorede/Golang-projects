package main

import (
	"fmt"
)

func main() {
	answer, err := Divide(6, 2)
	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
		return
	}

	fmt.Printf("%f\n", answer)

	total := Sum(5.5, 15.0, 24.6, 4.4)
	fmt.Printf("Total sum is: %f\n", total)
}
