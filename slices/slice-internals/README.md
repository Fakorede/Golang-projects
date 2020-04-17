# Slice Internals

### Backing Arrays

A slice doesn't directly store any elements.

Slice literals creates a hidden array and returns a slice that refers to that array. This hidden array is the _backing array_.

The _backing array_ stores the elements and is stored seperately from the slice, so it can be shared across multiple slices.

A slice is a window to its _backing array_.

A slice expression does not create a new array but rather creates a new slice value referring to some parts of the same _backing array_.The _backing array_ of the original slice is shared.

Slices implements changes in the _backing array_.

Indexing a slice is fast because the slice's _backing array_ is contiguous in memory.

A non-empty slice literal always creates a new _backing array_.

### Slice Header

A _Slice Value_ refers to its backing array through a memory address.

A slice needs to know the memory location of its array in order to find it on the computer memory. Go implements a data structure called _Slice Header_ which stores the memory address of a slice's backing array.

The _Slice Header_ has three(3) fields: `Pointer`, `Length`, `Capacity`.

`Pointer` stores the memory location of the slice's backing array.
`Length` is the no of elements the array contains.
`Capacity` stores how much space a backing array has.

**NOTE**

- A nil slice doesn't have a backing array but it has a slice header.

- Slice operations are cheap.

- Slicing creates a new slice header.

- Assigning a slice to another slice or passing it to a function only copies the slice header.

- Slice header has a fixed size and doesn't change even if we've millions of elements.

- On the other hand, an Array can be expensive. Assigning it or passing it to a function copies all its elements.

### Capacity of a Slice

```
ages := []int{10, 20, 30}
```

The slice literal has 3 elements, so it creates a backing array with 3 elements as well.

The _Capacity_ is the length of the backing array of the slice(not the slice). A slice created by a slice literal will have the same length and capacity.

The Capacity determines how much we can extend a slice when slicing/reslicing.

### Mechanics of Append

When the capacity is full, `append` allocates a new and larger backing array. It creates and returns a new slice header that points to the newly allocated array. After allocating the new array, it copies the elements of the previous array to the new array. And lastly, it assigns zero value to the uninitialized elements of the backing array.

```
ages := []int{10, 20}
```

The slice header has a length of 2 and a capacity of 2.

```
ages := []int{10, 20}
ages = append(ages, 30)
```

When we `append` another element to it, a new backing array is created with a length of 3 and capacity of 4. This is so as to prevent several allocations anytime `append` is called.

### Full Slice Expressions

Allows us to control the capacity of a returned slice from a slice expression.

```
newSlice := sliceable[START:STOP] // slice expression
newSlice := sliceable[START:STOP:CAP] // full slice expression
```

`CAP` limits the capacity of a returned slice from the slice expression.

```
sliceable := []byte{'f', 'u', 'l', 'l'}

fmt.Printf("%v\n", cap(sliceable[0:3])) // 4
fmt.Printf("%v\n", cap(sliceable[0:3:3])) // 3
```

### make

The `make()` function allows us pre-allocate backing array of a slice with a given length and capacity.

```
make([]int, 3) // [0 0 0] len - 3, cap - 3
```

`make()` initializes and returns a slice with the given length/capacity.
`[]int` is the type of the slice.
`3` reps the length and capacity of the slice.

We can also create a slice with different length and capacity. This is needed when we want to use only some parts of a slice.

```
make([]int, 3, 5) // len - 3, cap - 5
```

when the length is 0, the `make()` will return a slice ready to append.

```
make([]int, 0, 5) // len - 0, cap - 5
```

### copy

`copy()` copies the elements of slice into another slice.

```
copy(slice1, slice2)
```

`slice1` is the destination slice.
`slice2` is the source slice.

The element types of the slices should be identical.

```
package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	evens := []int{2, 4}
	odds := []int{3, 5, 7}

	s.Show("evens [before]", evens)
	s.Show("odds  [before]", odds)

	N := copy(evens, odds)
	fmt.Printf("%d element(s) are copied.\n", N) // 2 elements are copied

	s.Show("evens [after]", evens) // even [after] {3, 5}
	s.Show("odds  [after]", odds) // odd [after] {3, 5, 7}
}

```

The `copy()` function returns the number of elements it copies.

### Multi-dimensional Slices

We declare a multi-dimensional slice like so:

```
spendings := [][]int{
    {200, 100},   // 1st day
    {500},        // 2nd day
    {50, 25, 75}, // 3rd day
}
```

We can also use the `make()` function to declare a multi-dimensional slice:

```
spendings := make([][]int, 0, 5)

spendings := append(spendings, []int{200, 100})
spendings := append(spendings, []int{25, 10, 45, 60})
spendings := append(spendings, []int{5, 105, 35})
spendings := append(spendings, []int{95, 10}, []int{50, 25})
```
