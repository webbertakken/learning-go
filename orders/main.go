package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 33
	name := "Webber"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world!\n")

	// Println
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier

	// Default format
	fmt.Printf("my age is %v and my name is %v \n", age, name)

	// Quoted
	fmt.Printf("my age is %q and my name is %q \n", strconv.Itoa(age), name)

	// Type
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("name is of type %T \n", name)

	// Float
	fmt.Printf("Floating point with five decimals %.5f \n", 1.25)

	// Save formatted string
	var someStatement = fmt.Sprintf("My age is %v and my name is %v", age, name)
	fmt.Println(someStatement)
}
