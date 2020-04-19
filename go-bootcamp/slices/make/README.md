# make

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

**NOTE:** We only use `make()` when we want to optimize our software or pass a limited slice size to some function.
