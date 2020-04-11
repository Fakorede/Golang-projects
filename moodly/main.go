package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var moods = [...]string{
	"bad 👎",
	"sad 😞",
	"terrible 😩",
	"happy 😀",
	"good 👍",
	"awesome 😎",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pass a name as argument.")
		return
	}

	name := os.Args[1]

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[n])
}
