package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	answer, err := divide(6, 2)
	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
		return
	}

	fmt.Printf("%f\n", answer)
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

func divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}

	return p1 / p2, nil
}
