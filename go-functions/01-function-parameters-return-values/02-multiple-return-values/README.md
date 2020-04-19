# Multiple Return Values

We can define a function with multiple return values in GO as in below

```
func divide(p1, p2 float64) (float64, error) {
```

This means our program will always return two values regardless of the outcome, either a `float64` value and a `nil error` if things go well or a `NaN` value and a `error` value.

We get both values in different variables when we call the function

```
answer, err := divide(6, 2)
```

If we get an error, we handle like so:

```
if err != nil {
    fmt.Printf("An error occured %s\n", err.Error())
    return
}
```

Or we print our result if things go well

```
fmt.Printf("%f\n", answer)
```
