//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favouriteColour = "magenta"
	fmt.Println("My favourite colour is", favouriteColour)

	dateOfBirth, age := "September 88", 33
	fmt.Println("I was born in", dateOfBirth+".", "That makes me", age, "years old.")

	var (
		firstInitial  = "W"
		secondInitial = "T"
	)
	fmt.Println("My initials of my nickname are", firstInitial+"."+secondInitial+".")

	var ageInDays int
	ageInDays = age * 365
	fmt.Println("My age in days is", ageInDays)
}
