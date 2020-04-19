The lines of code below:

```
type MathExp = string

const (
	addExp = MathExp("add")
	subExp = MathExp("subtract")
	mulExp = MathExp("multiply")
	divExp = MathExp("divide")
)
```

Is simply a refactor of:

```
const addExp = "add"
const subExp = "subtract"
const mulExp = "multipy"
const divExp = "divide"
```

We need the codes above to be able to compare with the cases in the switch statement of the `mathExpression` function:

```
func mathExpression(exp MathExp) func(float64, float64) float64 {
	switch exp {
	case addExp:
		return simplemath.Add
	case subExp:
	 	return simplemath.Subtract
	case mulExp:
		return simplemath.Multiply
	case divExp:
	 	return simplemath.Divide
	default:
		return func(f1 float64, f2 float64) float64 {
			return 0
		}
	}
}
```

These codes simply return a function which takes 2 float64's and returns a float64 value based on the case.

In the `fmt.Printf` function of the `main` function, we call the `double` function. as an parameter along with two numbers and performs specified maths operations.

```
func main() {
	fmt.Printf("%f\n", double(3, 2, mathExpression(addExp)))
	fmt.Printf("%f\n", double(3, 2, mathExpression(mulExp)))
}
```

The `double` function takes the `mathExpression` as parameter alongside two float64 numbers.
The `double` function doubles the outcome of the specified mathematical expression in the `mathExp` and then returns the result.

```
func double(f1, f2 float64, mathExp func(float64, float64) float64) float64 {
	return 2.0 * mathExp(f1, f2)
}
```
