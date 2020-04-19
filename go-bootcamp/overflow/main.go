package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		width  uint8 = 255
		height       = 255
	)

	width++ // 256? the highest value a byte(uint8) can rep is 255. So it resets to its min value 0

	if int(width) < height {
		fmt.Println("height is greater")
	}

	fmt.Printf("width: %d height: %d\n", width, height) // width: 0, height: 255

	var n int8 = math.MaxInt8

	fmt.Println("max int8	:", n)       // 127
	fmt.Println("max int8 + 1	:", n+1) // -128, 127 + 1 = 128(overflow) but we get -128. This is because Go doesn't overflow. Instead, it wrap-arounds!

	f := float32(math.MaxFloat32)
	fmt.Println("Max Float	:", f*1.2)

}
