//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func sumFiles(files []string) {
	totalSum := 0
	sumFile := func(fileName string) {
		fileSum := 0
		// Note: you must run this from zero-to-mastery/lectures/exercise/goroutines for the file to be opened correctly.
		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("Error opening file: %v, %v", fileName, err)
		} else {
			log.Printf("%v: file opened.\n", file.Name())
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Printf("Error closing file:", err)
			} else {
				log.Printf("%v: file closed.\n", file.Name())
			}
		}(file)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() == true {
			if number, err := strconv.Atoi(scanner.Text()); err == nil {
				fileSum += number
			} else {
				log.Printf("%v: Can't parse line: %v, %v\n", file.Name(), scanner.Text(), err)
			}
		}

		log.Printf("%v: sum is %v\n", fileName, fileSum)
		totalSum += fileSum
	}

	for _, fileName := range files {
		go sumFile(fileName)
	}

	time.Sleep(100 * time.Millisecond)
	log.Printf("Total: %v\n", totalSum)
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	sumFiles(files)
}
