package main

import "fmt"

func main() {
	fmt.Printf("%f\n", add(2, 8))
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}
