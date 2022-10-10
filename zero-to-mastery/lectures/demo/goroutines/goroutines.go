package main

import (
	"fmt"
	"log"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}

	var capitalised []rune

	capitalise := func(r rune) {
		capitalisedCharacter := unicode.ToUpper(r)
		capitalised = append(capitalised, capitalisedCharacter)
		log.Printf("%c", capitalisedCharacter)
	}

	// Note that like this, write access to `capitalised` variable is not mutually exclusive and CHARACTERS GO MISSING.
	// Later in the course, synchronization with Mutex, communication via channels, and WaitGroups are all discussed which
	// solve the issue encountered in this demo.

	for _, character := range data {
		go capitalise(character)
	}

	for i := 0; i < len(data); i++ {
		go capitalise(data[i])
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Capitalised: %c\n", capitalised)
}
