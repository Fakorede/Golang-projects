package main

import (
	"fmt"

	"github.com/fakorede/learning-golang/go-functions/03-function-and-method-expressions/03-functions-as-parameters/simplemath"
)

// MathExp : not exported
type MathExp = string

const (
	addExp = MathExp("add")
	subExp = MathExp("subtract")
	mulExp = MathExp("multiply")
	divExp = MathExp("divide")
)

// const addExp = "add"
// const mulExp = "multipy"

func main() {
	fmt.Printf("%f\n", double(3, 2, mathExpression(addExp)))
	fmt.Printf("%f\n", double(3, 2, mathExpression(mulExp)))
}

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

func double(f1, f2 float64, mathExp func(float64, float64) float64) float64 {
	return 2.0 * mathExp(f1, f2)
}
