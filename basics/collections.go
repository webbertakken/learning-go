package main

import "fmt"

/*
An Array has a fixed length. It has to be specified during initialisation.

Slices are more like javascript-arrays and do not specify length.
Adding items to a Slice will create a new slice from scratch
*/

func collections() {
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

	// Maps
	itemPrices := map[string]float64{
		"soup":           4.19,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	for item, price := range itemPrices {
		fmt.Println(item, "-", price)
	}

	phonebook := map[int]string{
		267584967: "Mario",
		984759373: "Luigi",
		845775485: "Peach",
	}

	fmt.Printf("The number %d belongs to %s.\n", 267584967, phonebook[267584967])

	phonebook[984759373] = "Bowser"
	fmt.Printf("The number %d now belongs to %s.\n", 984759373, phonebook[984759373])

}
