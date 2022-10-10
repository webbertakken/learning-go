package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	defer wg.Done()

	wait()

	hits.Lock()
	defer hits.Unlock()

	hits.count += 1
	fmt.Printf("%04d: done\n", iteration)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	group := sync.WaitGroup{}
	hitCounter := Hits{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go serve(&group, &hitCounter, i)
	}

	group.Wait()
	fmt.Printf("Total hits: %v", hitCounter.count)
}
