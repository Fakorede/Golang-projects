package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jdoe@gmail.com",
			zipCode: 4200,
		},
	}

	john.updateName("Jane")
	john.print()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
