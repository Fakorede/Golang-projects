package main

import "fmt"

func main() {
	func() {
		fmt.Println("My first anonymous function...")
	}()

	value := func() {
		fmt.Println("Welcome to reality!")
	}

	value()
}
