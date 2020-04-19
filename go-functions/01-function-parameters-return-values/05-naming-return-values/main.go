package main

import (
	"errors"
	"fmt"
)

func main() {
	answer, err := divide(6, 2)
	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
		return
	}

	fmt.Printf("%f\n", answer)

	total := sum(5.5, 15.0, 24.6, 4.4)
	fmt.Printf("Total sum is: %f\n", total)
}

func sum(values ...float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}

	return total
}

func divide(p1, p2 float64) (ans float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}

	ans = p1 / p2

	return
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}
