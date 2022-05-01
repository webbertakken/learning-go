package main

import "fmt"

func inference() {
	// There are two syntaxes for inferring types

	// Using var
	var favouriteSnack = "ice cream" // string

	// Using := notation without var, cannot be used outside functions
	favouriteFlavour := "mango" // string

	// Output
	fmt.Println("My favorite snack is " + favouriteSnack + " with " + favouriteFlavour + " flavour")
}
