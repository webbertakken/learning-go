package main

import (
	"fmt"
	"math"
	"strings"
)

func functions() {
	fullNames := []string{"huey duck", "dewey duck", "louie duck"}

	// passing functions
	cycleNames(fullNames, sayGreeting)
	cycleNames(fullNames, sayBye)

	// return types
	radius := 12.0
	area := circleArea(radius)
	fmt.Printf("The area of a circle with radius %.1f is %.2f.\n", radius, area)

	// multiple return values
	for _, fullName := range fullNames {
		firstInitial, secondInitial := getInitials(fullName)
		fmt.Println(firstInitial, secondInitial)
	}
}

func sayGreeting(name string) {
	fmt.Printf("Good morning %s.\n", name)
}
func sayBye(name string) {
	fmt.Printf("Good bye %s.\n", name)
}

func cycleNames(names []string, fn func(string)) {
	for _, name := range names {
		fn(name)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func getInitials(fullName string) (string, string) {
	upper := strings.ToUpper(fullName)
	names := strings.Split(upper, " ")

	var initials []string
	for _, name := range names {
		initials = append(initials, name[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[len(initials)-1]
	}

	return initials[0], "_"
}
