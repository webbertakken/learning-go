package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	// Switches don't fall through by default.
	// For that you would need the `fallthrough` keyword
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap man")
	case p < 10:
		fmt.Println("moderately priced")
	default:
		fmt.Println("quite expensive!")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First class seating")
	default:
		fmt.Println("Other seating")
	}
}
