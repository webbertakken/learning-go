package main

import "fmt"

func main() {

	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println("Hello "+nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Toad"

	fmt.Println("Hello "+nameOne, nameTwo, nameThree)

	nameFour := "Yoshi" // shorthand, can not be used outside a function

	fmt.Println("And", nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 255

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8238492394923497239547252.18
	scoreThree := 1.5 // float 64

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
