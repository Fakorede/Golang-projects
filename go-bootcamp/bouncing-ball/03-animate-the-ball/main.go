package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width  = 50
		height = 10

		emptyCell = ' '
		ballCell  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		px, py int
		vx, vy = 1, 1

		cell rune // int32
	)

	// create multi-dimensional array
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := 0; y < height; y++ { // vertical
			for x := 0; x < width; x++ { // horizontal
				board[x][y] = false
			}
		}

		// ball position
		board[px][py] = true

		buf := make([]rune, 0, width*height)
		buf = buf[:0]

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

		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)

	}
}
