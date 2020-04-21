# Composition

Go gives us the ability to embed interfaces into interfaces, interfaces into structs, and structs into other structs.

### Embedding structs and Acessing Embedded fields

`Employee` and `Person` struct has fields which are common. We can thereby define a new `Name` struct which holds these fields and embed the `Name` struct in the `Employee` and `Person` structs.

We can also define a `FullName` method on the `Name` struct which the `Employee` and `Person` struct has access to.

```
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
}

// NewPerson : returns the Person object
func NewPerson(firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}
```

### Embedding interfaces

We can embed an `interface` into a `struct` type. There is no field on an interface, we're embedding the method sets of the interface on the struct type.

```
// Identifiable : Identifiable interface
type Identifiable interface {
	ID() string
}

// Person : Person struct
type Person struct {
	Name
	twitterHandler TwitterHandler
	Identifiable
}
```

We can also embed an `interface` into a `interface`

```
// Identifiable : Identifiable interface
type Identifiable interface {
	ID() string
}

// Citizen : Citizen interface
type Citizen interface {
	Identifiable
	Country() string
}
```
