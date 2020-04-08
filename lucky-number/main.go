package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, the harder it gets.

Wanna play?
`
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if turn == 0 && n == guess {
			fmt.Println("🙌	You are a genius. YOU WIN on your first guess!")
			return
		}

		if n == guess {
			fmt.Println("🎉	YOU WIN!")
			return
		}

		// fmt.Println(n)

	}

	fmt.Println("👻	YOU LOST...Try again?")
}
