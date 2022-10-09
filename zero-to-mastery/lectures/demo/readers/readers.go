package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0
	for {
		input, inputErr := r.ReadString('\n')
		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin", inputErr)
		}

		n := strings.Trim(input, " \r\n")
		if n == "" {
			continue
		}

		num, convertErr := strconv.Atoi(n)
		if convertErr != nil {
			fmt.Printf("Error converting \"%v\" to a number\n", n)
			continue
		}

		sum += num
		fmt.Println("sum:", sum)
		fmt.Print("input: ")
	}
}
