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

// Citizen : Citizen interface
type Citizen interface {
	Identifiable
	Country() string
}

// SocialSecurityNumber : SocialSecurityNumber
type socialSecurityNumber string

// NewSocialSecurityNumber : socialSecurityNumber constructor
func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

// europeanUnionIdentifier : europeanUnionIdentifier
type europeanUnionIdentifier struct {
	id      string
	country string
}

// NewEuropeanUnionIdentifier : europeanUnionIdentifier constructor
func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eu europeanUnionIdentifier) ID() string {
	return eu.id
}

func (eu europeanUnionIdentifier) Country() string {
	return eu.country
}

// Name struct : Name struct
type Name struct {
	first string
	last  string
}

// FullName : returns the first and last fields of the Name struct
func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

// Employee : Employee struct
type Employee struct {
	Name
}

// Person : Person struct
type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}

// NewPerson : Person constructor
func NewPerson(firstName string, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

// ID : method of the Identifiable interface implemented by Person
func (p *Person) ID() string {
	return p.Citizen.ID()
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
func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
