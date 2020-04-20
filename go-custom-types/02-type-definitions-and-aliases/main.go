package main

import (
	"fmt"

	"github.com/fakorede/learning-golang/go-custom-types/01-interfaces-and-structs/organization"
)

func main() {
	//var p organization.Identifiable = organization.Person{FirstName: "John", LastName: "Doe"}
	p := organization.NewPerson("John", "Doe")
	err := p.SetTwitterHandler("@jdoe")
	if err != nil {
		fmt.Printf("An error occurred setting Twitter handler: %s\n", err.Error())
	}

	fmt.Println(p.ID())
	fmt.Println(p.FullName())

	h := p.TwitterHandler()
	fmt.Println(h)
	// fmt.Println(h.RedirectURL())
}
