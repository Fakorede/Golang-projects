# Multi-dimensional Slices

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
