# copy

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
