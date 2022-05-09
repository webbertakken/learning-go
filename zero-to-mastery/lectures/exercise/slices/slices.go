//--Summary:
//  Create a program to manage parts on an assembly line.

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func printLine(title string, line []Part) {
	fmt.Println("\n---", title, "---")
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
}

func main() {
	//* Perform the following:
	//* Using a slice, create an assembly line that contains type Part
	//  - Create an assembly line having any three parts
	line := []Part{"Metal ore", "Copper ore", "Coal"}
	printLine("Start", line)

	//  - Add two new parts to the line
	line = append(line, "Metal plate", "Copper plate")
	printLine("After producing components", line)

	//  - Slice the assembly line, so it contains only the two new parts
	line = line[3:]
	printLine("When ores are used up", line)
}
