package main

import "fmt"

func loops() {
	// Loops
	names := []string{"Mario", "Luigi", "Yoshi", "Peach"}

	whileLoop()
	forLoop()
	forLoopIteratingASlice(names)
	forOfRangeLoop(names)
	usingContinueAndBreak(names)
}

func whileLoop() {
	x := 0
	for x < 5 {
		fmt.Println("The value of x is", x)
		x += 1
	}
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("The value of i is", i)
	}
}

func forLoopIteratingASlice(names []string) {
	for i := 0; i < len(names); i++ {
		fmt.Printf("The value at index %d is %s\n", i, names[i])
	}
}

func forOfRangeLoop(names []string) {

	for index, value := range names {
		fmt.Printf("The value at index %d is %s\n", index, value)

		// this only updates the local variable, not the original slice
		value = "something else"
	}

	// Not using index
	for _, value := range names {
		fmt.Printf("The value is still %s\n", value)
	}
}

func usingContinueAndBreak(names []string) {
	for index, value := range names {
		if index == 1 {
			fmt.Println("skipping at position", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at index", index)
			break
		}

		fmt.Printf("The value at position %d is %v\n", index, value)
	}
}
