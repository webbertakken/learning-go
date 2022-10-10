package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}

	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter += 1

		go func(i int) {
			defer func() {
				fmt.Printf("%v groups left to complete\n", counter)
				counter -= 1
				wg.Done()
			}()

			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Printf("%02d: Waiting for %v\n", i, duration)
			time.Sleep(duration)
		}(i)
	}

	fmt.Printf("Waiting for goroutines to all finish.\n")
	wg.Wait()
	fmt.Printf("Done!\n")
}
