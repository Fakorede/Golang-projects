package main

import (
	"fmt"
	"strings"
)

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {
	songs := []song{
		{title: "Intentions", artist: "Justin Bieber"},
		{title: "Yummy", artist: "Justin Bieber"},
	}

	rock := playlist{genre: "hip-hop", songs: songs}

	rock.songs[1].title = "Boyfriend"

	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	fmt.Printf(strings.Repeat("-", 42))
	for _, s := range rock.songs {
		fmt.Printf("\n%-20s %20s", s.title, s.artist)
	}

}
