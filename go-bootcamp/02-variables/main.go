package main

import (
	"fmt"
)

func main() {
	// ZERO VALUES
	var speed int
	var temperature float64
	var off bool
	var name string

	fmt.Println(speed)       // 0
	fmt.Println(temperature) // 0
	fmt.Println(off)         // false
	fmt.Printf("%q\n", name) // ""

	// MULTIPLE VARIABLE DECLARATION
	/*
		var (
			speed       int
			temperature float64
			off         bool
			name        string
		)
	*/

	// TYPE INFERENCE
	var age = 10
	fmt.Printf("Variable has a type %T\n", age)

	// SHORT DECLARATION
	isSafe := true
	fmt.Printf("Variable has a type %T\n", isSafe)

	// MULTIPLE SHORT DECLARATION
	month, day := "April", 11
	fmt.Println(month, day)

	// REDECLARATION
	var firstname string
	firstname, lastname := "Nikola", "Tesla"

	fmt.Println(firstname, lastname)
}
