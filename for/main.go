package main

import "fmt"

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}

		sum += i
		i++
	}

	fmt.Println(sum)
}
