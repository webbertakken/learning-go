//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"
)

//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

type Counter struct {
	sync.Mutex
	letters int
}

func (c *Counter) Add(amount int) {
	c.Lock()
	defer c.Unlock()
	c.letters += amount
}

func (c *Counter) Results() {
	log.Printf("%v letters\n", c.letters)
}

func (c *Counter) Reset() {
	c.letters = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value := scanner.Text()
		words := strings.Split(value, " ")

		counter := Counter{}
		group := sync.WaitGroup{}
		for index, word := range words {
			group.Add(1)
			go analyseWord(index, &group, &counter, word)
		}

		group.Wait()
		counter.Results()
		counter.Reset()
	}
}

func analyseWord(index int, group *sync.WaitGroup, counter *Counter, word string) {
	defer group.Done()

	count := 0
	for _, character := range word {
		if unicode.IsLetter(character) {
			count += 1
		}
	}

	counter.Add(count)
}
