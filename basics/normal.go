package main

import (
	"fmt"
	"reflect"
)

func normalSyntaxInGo() {
	// Initialise and infer multiple variables on one line
	amount, unit := 10, "tonnes"
	fmt.Println(amount, "("+reflect.TypeOf(amount).String()+")", unit, "("+reflect.TypeOf(unit).String()+")")

	// String concatenation
	var commentThatWasntAskedFor = "That's" + " a lot"
	commentThatWasntAskedFor += "!"
	fmt.Println(commentThatWasntAskedFor)
}
