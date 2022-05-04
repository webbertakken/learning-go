//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func printPersonsName(name string) {
	fmt.Println("Persons name is", name)
}

func randomMessage(length int) string {
	words := strings.Fields("Find box a little too small and curl up with fur hanging out . Stand in doorway, unwilling to chose whether to stay in or go out. Meow. Attack like a vicious monster refuse to drink water except out of someone's glass chill on the couch table for fall asleep upside-down. A nice warm laptop for me to sit on fall asleep upside-down need to check on human, have not seen in an hour might be dead oh look, human is alive, hiss at human, feed me where is it? i saw that bird i need to bring it home to mommy squirrel! and run off table persian cat jump eat fish gnaw the corn cob.")

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return strings.TrimSpace(strings.Join(words, " ")[:length])
}

func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func randomNumber() int {
	return int(rand.Int31())
}

func twoRandomNumbers() (int, int) {
	return int(rand.Int31()), int(rand.Int31())
}

func main() {
	// Need to seed random from the random function (only once)
	rand.Seed(time.Now().UnixNano())

	printPersonsName("Jason")

	fmt.Println("Some random words here you go: " + randomMessage(50))
	fmt.Println("1, 2 and 3 added make", addThreeNumbers(1, 2, 3))
	firstNumber := randomNumber()
	fmt.Println("Today's random number is", firstNumber)

	secondNumber, thirdNumber := twoRandomNumbers()
	fmt.Println("Two random numbers: ", secondNumber, thirdNumber)

	fmt.Println("Three random numbers added", addThreeNumbers(firstNumber, secondNumber, thirdNumber))
}
