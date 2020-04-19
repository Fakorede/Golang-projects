# Variadic Functions

They take multiple values in form of a single paramter passed into a function.

```
func sum(values ...float64) float64 {
    // block of code
}
```

To do anything with this values such as addition, we have to loop these values as they are actually seen by Go as a slice, in this case `[]float64`.

```
    for _, value := range values {
        // block of code
    }
```

To make it more obvious that we're passing a slice of `[]float64` values to the function, we can re-define the below line:

```
total := sum(5.5, 15.0, 24.6, 4.4)
```

as:

```
numbers := []float{5.5, 15.0, 24.6, 4.4}
total := sum(numbers...)
```

**Note:** Variadic parameters need to be passed after other parameters.

Defining as such will cause the Go compiler to raise an error:

```
func sum(values ...float64, price float64) float64 {
```
