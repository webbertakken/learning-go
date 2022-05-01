package main

import "fmt"

func inference() {
	// There are two syntaxes for inferring types

	// Using var
	var favouriteSnack = "ice cream" // string

	// Using := notation without var
	favouriteFlavour := "mango" // string

	// Output
	fmt.Println("My favorite snack is " + favouriteSnack + " with " + favouriteFlavour + " flavour")
}
