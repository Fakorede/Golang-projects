package main

import "fmt"

// Function
func entry(lang *string, aname *string) {

	// When the value of lang is nil it will panic
	if lang == nil {
		panic("Error: Language cannot be nil")
	}

	// When the value of aname is nil it will panic
	if aname == nil {
		panic("Error: Author name cannot be nil")
	}

	// When the values of the lang and aname are non-nil values it will print normal output
	fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
}

func main() {

	alang := "GO Language"

	// Here in entry function, we pass  a non-nil and nil values
	// Due to nil value this method panics
	entry(&alang, nil)
}
