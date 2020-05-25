package main

import "fmt"

func main() {
	type person struct {
		name, lastname string
		age            int
	}

	var dicaprio person
	var picasso person

	dicaprio.name = "Leornado"
	dicaprio.lastname = "DiCaprio"
	dicaprio.age = 47

	picasso.name = "Pablo"
	picasso.lastname = "Picasso"
	picasso.age = 91

	fmt.Printf("\nDiCaprio: %#v", dicaprio)
	fmt.Printf("\nPicasso: %#v", picasso)

	fmt.Printf("\n\n%s's age is %d", picasso.lastname, picasso.age)

}
