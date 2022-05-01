package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// Create bill
	name, _ := getInput("Guest name: ", reader)
	bill := newBill(name)

	return bill
}

func promptOptions(bill bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add item, t - add tip, s - save bill):", reader)
	switch option {
	case "a":
		promptAdd(bill, reader)
	case "t":
		promptTip(bill, reader)
	case "s":
		promptSave(bill)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(bill)
	}
}

func promptAdd(bill bill, reader *bufio.Reader) {
	itemName, _ := getInput("Item name:", reader)
	priceString, _ := getInput("Price:", reader)
	price, err := strconv.ParseFloat(priceString, 64)

	if err != nil {
		fmt.Println("The price must be a number")
		promptAdd(bill, reader)
	}

	bill.addItem(itemName, price)
	fmt.Printf("Added %s (€%.2f)\n", itemName, price)
	promptOptions(bill)
}

func promptTip(bill bill, reader *bufio.Reader) {
	tipAmountString, _ := getInput("Enter tip amount (€):", reader)
	tipAmount, err := strconv.ParseFloat(tipAmountString, 64)

	if err != nil {
		fmt.Println("The tip must be a number, use 0 for no tip")
		promptTip(bill, reader)
	}

	bill.updateTip(tipAmount)
	fmt.Printf("Added the tip of €%.2f\n", tipAmount)
	promptOptions(bill)
}

func promptSave(bill bill) {
	fmt.Println(bill.format())
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(strings.TrimSpace(prompt) + " ")
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func main() {
	bill := createBill()
	promptOptions(bill)
}
