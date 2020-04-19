# Returning Functions From Functions

```
func main() {
	add := mathExpression()
	fmt.Println(add(5.0, 15.0))
}

func mathExpression() func(float64, float64) float64 {
	return func(f1 float64, f2 float64) float64 {
		return f1 + f2
	}
}
```

`mathExpression` returns a function which takes in 2 float64's and returns a float64.
