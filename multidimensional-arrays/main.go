package main

import (
	"fmt"
)

func main() {

	// student1 := [3]float64{5, 6, 1}
	// student2 := [3]float64{9, 8, 4}
	
	// var sum float64
	
	// sum += student1[0] + student1[1] + student1[2]
	// sum += student2[0] + student2[1] + student2[2]
	
	
	// const N = float64(len(student1) * 2)
	
	// fmt.Printf("Average Grade: %g\n", sum/N)

	/*
	*---------------------------------------------------
	*	 MULTI-DIMENSIONAL ARRAY
	*---------------------------------------------------
	*/

	students := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64

	// sum += students[0][0] + students[0][1] + students[0][2]
	// sum += students[1][0] + students[1][1] + students[1][2]
	
	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	const N = float64(len(students) * len(students[0]))

	fmt.Printf("Average Grade: %g\n", sum/N)
	
	
}
