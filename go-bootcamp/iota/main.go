package main

import "fmt"

func main() {
	const (
		sunday = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)
}
