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

	// #1 Uninitialized(nil) map: Read-only
	// var dict map[string]string

	// fmt.Printf("Zero value: %#v\n", dict)
	// fmt.Printf("No of keys: %d\n", len(dict))

	// #2: Map retrieval is O(1) — on average.
	// key := "good"
	// value := dict[key]
	// fmt.Printf("%q means %#v\n", key, value)

	// map literal
	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	dict["up"] = "yukarı"  // adds a new pair
	dict["down"] = "aşağı" // adds a new pair
	dict["good"] = "iyi"   // overwrites the value at the key: "good"
	dict["mistake"] = ""   // a key with a zero-value

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)

}
