package main

import (
	"fmt"

	"github.com/fakorede/learning-golang/go-functions/02-method-receivers-on-struct-and-scalar-types/01-method-declarations/simplemath"
)

func main() {
	sv := simplemath.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	fmt.Println(sv.String())
}
