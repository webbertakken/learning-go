//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	age := rand.Int31() % 107

	fmt.Printf("At %d years old, you're ", age)
	switch {
	case age <= 0:
		fmt.Println("a newborn")
	case age <= 3:
		fmt.Println("a toddler")
	case age <= 12:
		fmt.Println("a child")
	case age <= 17:
		fmt.Println("a teenager")
	default:
		fmt.Println("an adult")
	}
}
