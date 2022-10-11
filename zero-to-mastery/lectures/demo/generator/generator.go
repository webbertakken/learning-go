package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

func generateRandIntn(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Generating random numbers using generateRandInt")
	randInt := generateRandInt(1, 100)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("Generating random numbers using generateRandIntn")
	randIntnRange := generateRandIntn(5, 1, 100)
	for number := range randIntnRange {
		fmt.Println(number)
	}

	fmt.Println("Generating random numbers using generateRandIntn again, checking if channel is still open")
	randIntnRange2 := generateRandIntn(5, 1, 100)
	for {
		number, open := <-randIntnRange2
		if !open {
			break
		}

		fmt.Println(number)
	}

}
