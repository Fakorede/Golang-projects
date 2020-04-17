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

### Append

A built-in function for adding a new element to a slice. Append returns a new slice.

```
append(slice, newElement)
```

**Example**

```
nums := []int{1, 2, 3}
nums = append(nums, 4, 5)
```

#### Elipsis

Allows us to append a slice to another slice.

```
var nums []int
num1:= []int{1, 2, 3, 4, 5}
num2:= []int{6, 7, 8, 9, 10}

nums = append(nums, num1...)
nums = append(nums, num2...)

fmt.Printf("%v\n", nums)
```

### Slice Expressions

- We can convert an Array to a Slice by using the slice expression `:`.

```
nums := [5]int{5, 4, 3, 2, 1}
fmt.Printf("%T\n", nums) // [5]int
fmt.Printf("%T", nums[:]) // []int
```

- To see more use-cases of the _slice expression_, please go through this [post](https://medium.com/golangspec/slice-expressions-in-go-963368c20765).

- A _sliceable value_ is any value that can be sliced using a slice expression. Slices, Arrays and Strings are sliceable. Slicing creates a new slice by cutting a sliceable.

```
newSlice := sliceable[start:stop]
```

The `start` position should be the index of an element in the sliceable. The `stop` position should be the position of the element we want to slice up to.

**Examples**

```
msg := []byte{'y', 'e', 'l', 'l', 'o', 'w'}
fmt.Printf("%s\n", msg[0:4]) // yell

```

`msg[0:6]` is equivalent to `msg[:]`

### Slicing vs Indexing

_Slicing_ returns a slice with the same type of the sliced slice.

_Indexing_ returns a single element with the type of the indexed slice's element type.

```
items := []string{
    "pacman", "mario", "tetris", "doom",
    "galaga", "frogger", "asteroids", "simcity",
    "metroid", "defender", "rayman", "tempest",
    "ultima",
}

fmt.Printf("slicing : %T %[1]q\n", items[2:3])

fmt.Printf("indexing: %T %[1]q\n", items[2])

```

### Reslicing

_Reslicing_ involves slicing another sliced slice.

```
items := []string{
    "pacman", "mario", "tetris", "doom",
    "galaga", "frogger", "asteroids", "simcity",
    "metroid", "defender", "rayman", "tempest",
    "ultima",
}

last4 := items[l-4:]
mid := last4[1:3]

```
