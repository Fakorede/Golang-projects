package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// empty := ""
	// dump(empty)

	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")

	for i := range hello {
		dump(hello[i : i+1])
	}

	// conversions allocate new backing arrays
	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]rune(hello)))
}

type StringHeader struct {
	// points to a backing array's item
	pointer uintptr // where it starts
	length  int     // where it ends
}

// dump prints the string header of a string value
func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
