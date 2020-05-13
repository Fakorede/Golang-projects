package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}

	query := args[0]

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
	}

	// turkish := dict // #1 - both maps sees the changes
	// turkish["good"] = "güzel"
	// dict["great"] = "kusursuz"

	delete(dict, "awesome") // deletes a given key

	// delete all keys and values
	//dict = nil - replaces the pointer
	for k := range dict {
		delete(dict, k)
	}

	turkish := make(map[string]string, len(dict))
	for k, v := range dict {
		turkish[v] = k
	}

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)

}
