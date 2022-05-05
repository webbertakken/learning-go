package main

import "fmt"

func fibonacci(maxNumber int) {
	fmt.Println("Sum is", 1)
	for i, j := 0, 1; i <= maxNumber; {
		sum := i + j

		fmt.Println("Sum is", sum)

		i = j
		j = sum
	}
}

func main() {
	fibonacci(900)

	sum := 55
	for sum > 10 {
		sum -= 5
		fmt.Println("Subtract 5. Sum is now", sum)
	}
}
