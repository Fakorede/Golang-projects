**Type Aliases** are a reference to other types.

```
// TwitterHandler : type alias for twitterHandler
type TwitterHandler = string
```

```
// Person : Person struct
type Person struct {
	...
	twitterHandler TwitterHandler
}
```

**Type Definition**

```
// TwitterHandler : type definition for twitterHandler
type TwitterHandler string
```

_Type Aliases_ copies the fields and the methods of a type to a new type,thereby becoming that exact type. We cannot extend and add new methods on them.

A _Type Definition_ it only copies the fields of a type(underlying type) over to a new type, it does not copy the methods. We can add new methods.
We have to use _Explicit Type Conversion_ so it can access the method of the underlying type.

```
// TwitterHandler : type def for twitterHandler
type TwitterHandler string

// RedirectURL : RedirectURL
func (th TwitterHandler) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://twitter.com/%s", cleanHandler)
}
```
