# Slices

A Slice is a composite type in GO. It is a collection of elements which are indexable and is of a dynamic length i.e not fixed.

We can't assign elements to a _nil slice_

```
var games []string
games[0] = "FIFA"
games[1] = "GTA"

fmt.Printf("%#v\n", games)
```

A slice can only be compared to a nil value.

```
    games := []string{"FIFA", "GTA"}
    newGames := []string{"PES", "COD", "WWE"}

    if games != nil {
        fmt.Println("it works!")
    }

    if games == newGames {
    // invalid operation: games == newGames (slice can only be compared to nil)
    }
```

We can only compare slices can looping through the range.

```
games := []string{"FIFA", "GTA"}
newGames := []string{"PES", "COD", "WWE"}

var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = " not "
		}
	}

	fmt.Printf("games and newGames are %s equal\n", ok)
```

Unlike an array, we can assign slices with different lengths to each other.

```
games := []string{"FIFA", "GTA"}
newGames := []string{"PES", "COD", "WWE"}

games = newGames
fmt.Printf("%#v", games)
```

We can set a slice to nil.

```
games := []string{"FIFA", "GTA"}
newGames := []string{"PES", "COD", "WWE"}

games = newGames
newGames = nil
fmt.Printf("%#v\n", games) // []string{"PES", "COD", "WWE"}
fmt.Printf("%#v\n", newGames) // []string(nil)
```

Instead of checking whether a slice is nil or not, we can check for the length of a slice.

```
games := []string{"FIFA", "GTA"}

if games != nil {
    fmt.Println("it works!")
}

if len(games) != 0 {
    fmt.Println("also works!")
}
```

### Nil Slices and Empty Slices

_Nil slices_ are unintialized slices while _Empty slices_ are initialized slices.

```
var z []int
y := []int{}

fmt.Printf("%#v\n", z) // []int(nil)
fmt.Printf("%#v\n", y) // []int{}
```

### Slices vs Arrays

| Slices                                | Arrays                                |
| ------------------------------------- | ------------------------------------- |
| Slices are dynamic                    | Arrays are fixed                      |
| Slices can grow and shrink in runtime | They cannot grow or shrink            |
| Zero value is nil                     | Zero value is zero valued elements    |
| Its length is not part of its type    | Its length is part of its type        |
| Example declaration: `var num []int`  | Example declaration: `var num [5]int` |

#### Similar Properties

- They can only contain the same type of elements.
- We can access and change the elements using an index value.
- The len, cap, append built-in functions work for both.
