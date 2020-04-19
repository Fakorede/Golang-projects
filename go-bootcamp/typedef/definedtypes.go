package main

import "fmt"

type gram float64
type ounce float64

func definedTypes() {
	var g gram = 1000
	var o ounce

	//o = g * 0.035 - type mismatch
	o = ounce(g) * 0.035

	fmt.Printf("%g grams is equivalent to %.2f ounces\n", g, o)
}
