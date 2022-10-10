//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//

package main

import (
	"fmt"
	"unicode"
)

//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification
func classifyLines(collection []string) {
	var lines, characters, letters, digits, spaces, symbols, punctuationMarks int

	characterClassifier := func(rune rune) {
		if unicode.IsLetter(rune) {
			letters += 1
		}
		if unicode.IsDigit(rune) {
			digits += 1
		}
		if unicode.IsSpace(rune) {
			spaces += 1
		}
		if unicode.IsSymbol(rune) {
			symbols += 1
		}
		if unicode.IsPunct(rune) {
			punctuationMarks += 1
		}
	}

	lineClassifier := func(line string) {
		for _, character := range line {
			characters += 1
			characterClassifier(character)
		}
	}

	for _, line := range collection {
		lines += 1
		lineClassifier(line)
	}

	fmt.Print("Input has ")
	fmt.Printf("%d lines, ", lines)
	fmt.Printf("%d words, ", lines)
	fmt.Printf("%d characters, ", characters)
	fmt.Printf("%d letters, ", letters)
	fmt.Printf("%d digits, ", digits)
	fmt.Printf("%d spaces, ", spaces)
	fmt.Printf("%d symbols ", symbols)
	fmt.Print("and ")
	fmt.Printf("%d punctuation marks", punctuationMarks)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	//* Create a single function to classifyLines over each line of text that is
	//  provided in main().
	//  - The function must return nothing and must execute a closure
	classifyLines(lines)
}
