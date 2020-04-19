package main

import "fmt"

func main() {
	add := mathExpression()
	fmt.Println(add(5.0, 15.0))
}

func mathExpression() func(float64, float64) float64 {
	return func(f1 float64, f2 float64) float64 {
		return f1 + f2
	}
}
