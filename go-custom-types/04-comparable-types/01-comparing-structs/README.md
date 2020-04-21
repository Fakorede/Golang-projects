# Comparable Types

## Comparing Structs

Types with the same underlying Struct type are comparable

```
func main() {
	name1 := Name{"John", "Doe"}
	name2 := Name{"Mike", "Harrison"}

	if name1 == name2 {
		fmt.Println("they are comparable!") // they are comparable!
	}
}

type Name Struct {
    first string
    last string
}
```

But we cannot compare types with different Underlying Struct types. This produces a compile-time error.

```

func main() {
	name1 := Name{"John", "Doe"}
	name3 := AnotherName{"John", "Doe"}


	if name1 == name3 {
		fmt.Println("not comparable!")
	}
}

// Name : Name struct
type Name struct {
	first string
	last  string
}

// AnotherName : AnotherName struct
type AnotherName struct {
	first string
	last  string
}

```
