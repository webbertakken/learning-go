//--Summary:
//  Create a function that can parse time strings into component values.
//

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components

type Time struct {
	hours   int
	minutes int
	seconds int
}

type TimeParsingError struct {
	msg   string
	input string
}

func (t *TimeParsingError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseFromString(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParsingError{"Incorrect format", input}
	}

	hours, err := strconv.Atoi(components[0])
	if err != nil {
		return Time{}, &TimeParsingError{err.Error(), input}
	}

	minutes, err := strconv.Atoi(components[1])
	if err != nil {
		return Time{}, &TimeParsingError{err.Error(), input}
	}

	seconds, err := strconv.Atoi(components[2])
	if err != nil {
		return Time{}, &TimeParsingError{err.Error(), input}
	}

	if hours < 0 || hours > 23 {
		return Time{}, &TimeParsingError{"Hours must be a number between 0 to 23, received", components[0]}
	}
	if minutes < 0 || minutes > 59 {
		return Time{}, &TimeParsingError{"Minutes must be a number between 0 to 59, received", components[1]}
	}
	if seconds < 0 || seconds > 59 {
		return Time{}, &TimeParsingError{"Seconds must be a number between 0 to 59, received", components[2]}
	}

	return Time{hours, minutes, seconds}, nil
}

//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

func main() {
	time, err := ParseFromString("14:07:33")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.hours, "hours")
	fmt.Println(time.minutes, "minutes")
	fmt.Println(time.seconds, "seconds")
}
