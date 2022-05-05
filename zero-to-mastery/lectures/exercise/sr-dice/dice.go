//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func diceRoll(sides int) int {
	return int(rand.Int31n(int32(sides)) + 1)
}

func comboRoll(dice int, sides int, times int) {
	for i := 0; i < times; i++ {

		// Had to use []any to keep it practical (trying to pass each element as a separate argument to fmt.Println().
		// It seemed a bit much to make a helper function just to convert a slice of integers to strings.
		// https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string

		// I believe go language could be improved, unless I misunderstand something fundamental.
		// Asked a question on discord, hopefully some experts on go will be able to say something smart about it.
		// https://discord.com/channels/423464391791476747/933759051190370314/971905425572261919
		var numbers []any

		for j := 0; j < dice; j++ {
			numbers = append(numbers, diceRoll(sides))
		}

		fmt.Println(numbers...)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	comboRoll(3, 6, 3)
}
