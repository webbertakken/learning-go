package main

import (
	"fmt"
	"math"
)

func functions() {
	names := []string{"visitor1", "visitor2"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	radius := 12.0
	area := circleArea(radius)
	fmt.Printf("The area of a circle with radius %.1f is %.2f.\n", radius, area)
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
