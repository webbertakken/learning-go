package main

import (
	"fmt"
	"math"
)

func mathematics() {
	// Arithmetic
	number := 3 // 3
	number -= 3 // 0
	number += 3 // 3
	number *= 3 // 9
	number /= 3 // 3

	// Modulus
	rest := number % 2 // 1

	// using "math"
	squareRootOfThree := math.Sqrt(3)

	// Using Sprintf to format number types (%s = string, %d = digit, %f = float or %v = default format for any type)
	// Note that Sprintf uses "reflect" which creates an additional object, so it's not the most efficient
	fmt.Println(fmt.Sprintf("rest = %d, square root of three = %f", rest, squareRootOfThree))
}
