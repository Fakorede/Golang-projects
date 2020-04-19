package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var moods = [...][3]string{
	{"happy 😀", "good 👍", "awesome 😎"},
	{"bad 👎", "sad 😞", "terrible 😩"},
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Pass [name] and [mood] as argument. [mood] can be 'positive' or 'negative'")
		return
	}

	name := args[0]
	mood := args[1]

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1 * len(moods[0]))

	if mood == "positive" {
		fmt.Printf("%s feels %s\n", name, moods[0][n])
	} else {
		fmt.Printf("%s feels %s\n", name, moods[1][n])
	}

}
