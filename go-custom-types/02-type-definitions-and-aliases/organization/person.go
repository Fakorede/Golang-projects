package organization

import (
	"errors"
	"fmt"
	"strings"
)

// Handler : Handler
type Handler struct {
	handle string
	name   string
}

// TwitterHandler : type def for twitterHandler
type TwitterHandler string

// RedirectURL : RedirectURL
func (th TwitterHandler) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://twitter.com/%s", cleanHandler)
}

// Identifiable : Identifiable interface
type Identifiable interface {
	ID() string
}

// Person : Person struct
type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

// NewPerson : returns the Person object
func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

// FullName : returns the firstName and lastName fields of the Person struct
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// ID : method of the Identifiable interface implemented by Person
func (p Person) ID() string {
	return "12345"
}

// SetTwitterHandler : sets the TwitterHandler field on the Person struct
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler

	return nil

}

// TwitterHandler : returns the TwitterHandler field on the Person struct
func (p Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
