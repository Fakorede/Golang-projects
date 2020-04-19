package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	// s.Show("0:4", items[0:4])
	// s.Show("4:8", items[4:8])
	// s.Show("8:12", items[8:12])
	// s.Show("12:13", items[12:13])

	const pageSize = 4

	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize

		if to > l {
			to = l
		}

		currentPage := items[from:to]
		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)

		s.Show(head, currentPage)
	}

}
