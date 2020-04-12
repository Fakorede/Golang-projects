# Named and Unnamed Types

With type definition, we can create a new type with a name. The syntax for creating a named type is `type typeName underlyingType`. Example:

```
type bookcase [3]int
bookcase{6, 9, 3} // named type, underlying type is [3]int
```

An unnamed composite type underlying type is itself.

```
[3]int{6, 9, 3} // unnamed type, underlying type is [3]int
```

### Example

```
type bookcase [3]int

	a := bookcase{6, 9, 3}
	b := [3]int{6, 9, 3}

	if a == b {
		fmt.Println(true)
	}
```

Named and Unnamed types are comparable if their underlying types are identical.
