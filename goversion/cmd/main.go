package main

import (
	"fmt"

	"github.com/fakorede/learning-golang/goversion"
)

func main() {
	fmt.Println("The version of go found on your machine is:", goversion.Version())
}
