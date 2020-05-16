package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	os.Stdin.Close() // closes Standard Input
	in := bufio.NewScanner(os.Stdin)

	var lines int

	in.Scan()

	fmt.Println("Scanned Text :", in.Text())
	fmt.Println("Scanned Text :", in.Bytes())

	for in.Scan() {
		lines++
		fmt.Println("Scanned Text :", in.Text())
	}

	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
