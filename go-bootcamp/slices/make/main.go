package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	// #1: assume that tasks can be a long list
	// ---------------------------------------------
	tasks := []string{"jump", "run", "read"}

	// #2: INEFFICIENT WAY - allocates 3 diff backing array each
	// time we append to the upTasks.
	// ---------------------------------------------
	// var upTasks []string
	// s.Show("upTasks", upTasks)

	// for _, task := range tasks {
	// 	upTasks = append(upTasks, strings.ToUpper(task))
	// 	s.Show("upTasks", upTasks)
	// }

	// #3: SO SO WAY: we create a hidden backing array with
	// a length of 3(len(tasks)) and return a slice that points
	//  to that array. But it appends to the back of the slice
	// len and adds more capacity and backing array.
	// ---------------------------------------------
	// upTasks := make([]string, len(tasks))
	// s.Show("upTasks", upTasks)

	// for _, task := range tasks {
	// 	upTasks = append(upTasks, strings.ToUpper(task))
	// 	s.Show("upTasks", upTasks)
	// }

	// #4: SO SO WAY 2: starts appending from the zeroth
	// index of the slice and doesn't create more backing arrays.
	// ---------------------------------------------
	// upTasks := make([]string, len(tasks))
	// s.Show("upTasks", upTasks)

	// for i, task := range tasks {
	// 	upTasks[i] = strings.ToUpper(task)
	// 	s.Show("upTasks", upTasks)
	// }

	// #5: allocates a new array
	// when we append again, it still creates another backing array.
	// upTasks = append(upTasks, "PLAY")
	// s.Show("upTasks", upTasks)

	// #6: SO SO WAY 3
	// ---------------------------------------------
	// upTasks := make([]string, 0, len(tasks))

	// #7
	// upTasks = upTasks[:cap(upTasks)]

	// #6
	// s.Show("upTasks", upTasks)

	// for i, task := range tasks {
	// 	upTasks[i] = strings.ToUpper(task)
	// 	s.Show("upTasks", upTasks)
	// }

	// #8: THE BEST WAY
	// doesn't allocate a new array and we dont need
	// to introduce an index variable.
	// ---------------------------------------------
	upTasks := make([]string, 0, len(tasks))

	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
}
