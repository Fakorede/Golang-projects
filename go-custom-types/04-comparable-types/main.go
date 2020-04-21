package main

import "fmt"

func main() {
	name1 := Name{"John", "Doe"}
	name2 := Name{"John", "Doe"}
	// name3 := AnotherName{"John", "Doe"}

	if name1 == name2 {
		fmt.Println("they are comparable!")
	}

	// if name1 == name3 {
	// 	fmt.Println("they are comparable!")
	// }
}

// Name : Name struct
type Name struct {
	first string
	last  string
}

// AnotherName : AnotherName struct
type AnotherName struct {
	first string
	last  string
}
