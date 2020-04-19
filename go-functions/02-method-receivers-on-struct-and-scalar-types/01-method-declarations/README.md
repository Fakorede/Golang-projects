# Method Declaration

**Methods** makes it easier to encapsulate fields within an object and control how state is modified. Methods are shorter and have more succint naming convention.

In Go, we can define a method on a struct and choose if the receiver (the object which the method is executed on) is of type Value or a Pointer.

We define a type `struct` named `SemanticVersion`

```
type SemanticVersion struct {
	major, minor, patch int
}

```

We then define a function `NewSemanticVersion` which takes 3 parameters of type `int` and then returns a type of `SemanticVersion`.

```
func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}
```

### Value Receiver

We can now define a `method` on this struct which we call `String`. This method takes a receiver(object which we call this method on) which finally returns a string.

```
func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}
```

The receiver used here is a **Value Receiver**. The **Value Receiver** makes a copy of the type and passes it to the method. The method stack now holds an equal object but at a different location on memory.

### Pointer Receiver

Lets imagine we need to make changes to the versions in this object. We need to define new new methods

```
// IncrementMajor : this method increments and prints the major version
func (sv SemanticVersion) IncrementMajor() {
	sv.major++
}
```

Now, if when we call this method, we observe no change at all

```
func main() {
	sv := simplemath.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	fmt.Println(sv.String())
}
```

```
go run main.go
1.2.3
```

This is because the _Value Receiver_ passes different copies of the `sv` object to each method.

To get the desired behavior here, we need to use a **Pointer based method receiver**. The **Pointer Receiver** passes the address of a type to the method. The method stack now has a reference to the original object.

```
// IncrementMajor : this method increments and prints the major version
func (sv *SemanticVersion) IncrementMajor() {
	sv.major++
}
```

Now, when we run the program again, we get the desired output like so:

```
go run main.go
2.2.3
```

### Conclusion

In Go, it is idiomatic to set all of our methods to pointer-based receivers.

**Pointer Receivers** are very useful when modifying state within a type.

It is also very useful when we have a huge type and we do not want to be copying all over the place.
