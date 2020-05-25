package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const (
	message = "Welcome to Fab's Game Store. We currently have %d games available.\n"
	usage   = `
	The available commands are: id, list, save, and quit.
	
	> id	 : to get a specific game with an id
	> list   : lists all the games
	> save	 : exports the games to a file
	> quit   : quits the program
	> decode : decodes the saved json formatted data
	`
)

type indexType map[int]game

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

var output []byte

func main() {

	games := []game{
		{item{1, "god of war", 50}, "action adventure"},
		{item{2, "x-com 2", 30}, "strategy"},
		{item{3, "minecraft", 20}, "sandbox"},
	}

	index := make(indexType)
	for _, g := range games {
		index[g.id] = g
	}

	fmt.Println()
	fmt.Printf(message, len(games))
	fmt.Println(usage)

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()

		if !in.Scan() {
			break
		}

		fmt.Println()

		switch in.Text() {
		case "quit":
			fmt.Println("Goodbye!")
			return
		case "list":
			getGames(games)
		case "id":
			fmt.Println("Please provide an id:")
			in.Scan()

			fmt.Println()

			input := in.Text()
			id(games, input, index)
		case "save":
			save(games)
		case "decode":
			decode()
		}
	}

}

func decode() {
	var decoded []jsonGame
	if err := json.Unmarshal([]byte(output), &decoded); err != nil {
		fmt.Println("An error occured: ", err)
		return
	}

	for _, dg := range decoded {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			dg.ID, dg.Name, "("+dg.Genre+")", dg.Price)

		fmt.Println()
	}
}

func save(games []game) {

	var allGames []jsonGame

	for _, g := range games {
		allGames = append(allGames, jsonGame{g.id, g.name, g.genre, g.price})
	}

	output, err := json.MarshalIndent(allGames, "", "\t")
	if err != nil {
		fmt.Println("An error occured: ", err)
		return
	}

	fmt.Println(string(output))
	return
}

func id(games []game, input string, index indexType) {
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid id")
		return
	}

	g, ok := index[num]

	if num > len(index) || !ok {
		fmt.Println("Sorry. We don't have that game at the moment.")
		return
	}

	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}

func getGames(games []game) {
	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.id, g.name, "("+g.genre+")", g.price)
	}
}
