package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence string = "lazy cat jumps again and again and again"
	words := strings.Fields(sentence)

	// for i := 0; i < len(words); i++ {
	// 	fmt.Printf("#%-2d: %q\n", i+1, words[i])
	// }

	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
}
