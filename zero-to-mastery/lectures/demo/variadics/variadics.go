package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)

	answer := sum(all...)
	fmt.Println(answer)
}
