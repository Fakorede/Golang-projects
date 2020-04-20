An **Interface** is a custom type that is used to specify a set of one or more method signatures. **Interfaces** are abstract, so you we're not allowed to create an instance of the interface.

But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires. Or in other words, the interface is a collection of methods as well as it is a custom type.

Below is the syntax for creating an **Interface**

```
// Identifiable : Identifiable interface
type Identifiable interface {

}
```

In an **Interface** we can have functions and we don't have to explicitly use the `func` keyword because its the only thing we can have in there.

Any type that implements the `Identifiable` interface is of type `Identifiable`

A **Struct** is a user-defined type that contains a collection of named fields/properties. It is used to group related data together to form a single unit. Any real-world entity that has a set of properties can be represented using a struct.

If you’re coming from an object-oriented background, you can think of a **struct** as a lightweight class that supports composition but not inheritance.

Below is the syntax for creating a **Struct**

```
// Person : Person struct
type Person struct {

}
```

`Person` implicitly inherits the `Identifiable` interface by implementing the `ID` method.

```
// ID : method of the Identifiable interface implemented by Person
func (p Person) ID() string {
	return "12345"
}
```

We can define **Private Fields** on the `Person` struct:

```
// Person : Person struct
type Person struct {
	firstName string
	lastName  string
}
```

We define a **Constructor** which sets the fields of the `Person` struct:

```
// NewPerson : returns the Person object
func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}
```

We define a `FullName` method on the `Person` struct to a formatted version of the fields accessible.

```
// FullName : returns the firstName and lastName fields of the Person struct
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}
```

To **Change the State of the Struct**, we have to implement a **setter**

```
// Person : Person struct
type Person struct {
	...
	twitterHandler string
}
```

```
// SetTwitterHandler : sets the TwitterHandler field on the Person struct
func (p Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler

	return nil

}

```

And then create a **getter** to return the data:

```
// TwitterHandler : returns the TwitterHandler field on the Person struct
func (p Person) TwitterHandler() string {
	return p.twitterHandler
}
```

**Note:** When changing state, we have to make use of a _Pointer receiver_ to track and make our data uniform:

```
// SetTwitterHandler : sets the TwitterHandler field on the Person struct
func (p *Person) SetTwitterHandler(handler string) error {
	...
}
```

In our `main.go` file, we can create a NewPerson access and access all the needed information:

```
p := organization.NewPerson("John", "Doe")
err := p.SetTwitterHandler("@jdoe")
if err != nil {
	fmt.Printf("An error occurred setting Twitter handler: %s\n", err.Error())
}

fmt.Println(p.ID()) // 12345
fmt.Println(p.FullName()) // John Doe
fmt.Println(p.TwitterHandler()) // @jdoe
```
