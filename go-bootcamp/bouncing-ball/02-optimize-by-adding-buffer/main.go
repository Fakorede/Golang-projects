package main

import "fmt"

func main() {
	const (
		width  = 50
		height = 10

		emptyCell = ' '
		ballCell  = '⚾'
	)

	var cell rune // int32

	// create multi-dimensional array
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// ball position
	board[0][0] = true

	buf := make([]rune, 0, width*height)

	// draw the board
	for y := 0; y < height; y++ { // vertical
		for x := 0; x < width; x++ { // horizontal
			cell = emptyCell
			if board[x][y] {
				cell = ballCell
			}
			// fmt.Print(string(cell), " ")
			buf = append(buf, cell, ' ')

		}
		// fmt.Println()
		buf = append(buf, '\n')
	}
	fmt.Print(string(buf))
}
