package main

import (
	"fmt"
	"math"
)

func main() {

	p2 := powerOfTwo()
	value := p2()

	fmt.Println(value) // 4

	value = p2()

	fmt.Println(value) // 9
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x++
		return int64(math.Pow(x, 2))
	}
}
