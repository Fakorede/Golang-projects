package main

import "fmt"

func aliasedTypes() {
	var (
		byteVal  byte
		uint8Val uint8
		runeVal  rune
		int32Val int32
	)

	byteVal = uint8Val // valid cos byte and uint8 are aliases
	runeVal = int32Val // valid, rune and int32 are aliases

	fmt.Println(byteVal, runeVal)
}
