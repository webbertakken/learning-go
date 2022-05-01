package main

import "fmt"

/*
An Array has a fixed length. It has to be specified during initialisation.

Slices are more like javascript-arrays and do not specify length.
Adding items to a Slice will create a new slice from scratch
*/

func arrays() {
	// Arrays
	ages := [3]int{20, 25, 30}
	names := [4]string{"Yoshi", "Mario", "Luigi", "Peach"}
	names[1] = "Bowser"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]  // (not inclusive), so index 1 and 2
	rangeTwo := names[2:]   // from index to end
	rangeThree := names[:3] // from index to end
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Koopa")
	fmt.Println(rangeOne)
}
