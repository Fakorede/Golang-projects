package main

import "fmt"

func underlyingTypes() {
	type (
		// Gram's underlying type is int
		Gram int
		// Kilogram's underlying type is int
		Kilogram Gram
		// Ton's underlying type is int
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d", salt, apples, truck)

	// salt = apples // wont work cos they are of different types.
	salt = Gram(apples) // type conversion works here cos they're of the same underlying type
}
