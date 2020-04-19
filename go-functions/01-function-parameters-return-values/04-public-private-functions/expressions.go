package main

import (
	"errors"
	"math"
)

// Sum : sums up slice values provided as argument.
func Sum(values ...float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}

	return total
}

// Divide : divides argument p1 by argument p2
func Divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}

	return p1 / p2, nil
}

// Add : adds the provided arguments
func Add(p1, p2 float64) float64 {
	return p1 + p2
}
