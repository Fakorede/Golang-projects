package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]

	celsius, _ := strconv.ParseFloat(args, 64)

	fahrenheit := celsius*1.8 + 32

	fmt.Printf("%g°C is equivalent to %g°F.🌡️\n", celsius, fahrenheit)
}
