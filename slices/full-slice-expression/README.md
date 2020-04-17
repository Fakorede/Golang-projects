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

Use _full slice expression_ to prevent other code from appending more elements to a slice's backing array.

```
nums := []int{1, 3, 2, 4} // {1, 3, 2, 4} len - 2, cap - 4, a backing array
odds := nums[:2]          // {1, 3} len - 2, cap - 4, same backing array
odds := nums[:2:2]        // {1, 3} len - 2, cap - 2, same backing array
odds = append(odds, 5, 7) // {1, 3, 5, 7} len - 4, cap - 4, another backing array
odds := append(nums[:2:2], 5, 7) - same as two lines above(24 & 25) combined.

even := append(nums[2:4], 6, 8) // {2, 4, 6, 8}, len - 4, cap - 4
```
