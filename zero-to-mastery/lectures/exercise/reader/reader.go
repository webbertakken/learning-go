//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:

//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input string

func (i *Input) string() string {
	return string(*i)
}

func (i *Input) lower() string {
	return strings.ToLower(i.string())
}

//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
func (i *Input) handleGreeting() {
	input := i.lower()
	greetings := map[string]struct{}{
		"hi":    {},
		"hello": {},
	}
	if _, ok := greetings[input]; ok {
		isCommand = true
		fmt.Println("Howdy!")
	}
}

//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
func (i *Input) handleExit() {
	if i.lower() == "q" {
		numCommands += 1
		isExiting = true

		fmt.Println("Bye.")
	}
}

func printStats() {
	fmt.Printf("--Stats--\ninputs: %d\ncommands: %d\n\n", numInputs, numCommands)
}

func printAskForInput() {
	fmt.Print("input: ")
}

var numInputs int
var numCommands int
var isCommand bool
var isExiting bool

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numInputs = 0
	numCommands = 0

	printAskForInput()
	for scanner.Scan() {
		numInputs += 1
		isCommand = false
		input := Input(strings.Trim(scanner.Text(), " \n\r"))

		input.handleGreeting()
		input.handleExit()

		if isCommand {
			numCommands += 1
		}

		if isExiting {
			printStats()
			os.Exit(0)
		}

		printAskForInput()
	}
}
