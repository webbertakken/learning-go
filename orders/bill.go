package main

import (
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

func (bill *bill) format() string {
	total := 0.00
	formattedBill := strings.Builder{}

	// Heading
	formattedBill.WriteString("Bill for " + bill.name + ".\n")

	// Add items
	for item, price := range bill.items {
		total += price
		formattedBill.WriteString(fmt.Sprintf("%-25v €%v\n", item, price))
	}

	// Tip
	total += bill.tip
	formattedBill.WriteString(fmt.Sprintf("%-25v %v\n", "-----", "-----"))
	formattedBill.WriteString(fmt.Sprintf("%-25v €%.2f\n", "Tip", bill.tip))

	// Add total
	formattedBill.WriteString(fmt.Sprintf("%-25v %v\n", "-----", "-----"))
	formattedBill.WriteString(fmt.Sprintf("%-25v €%.2f\n", "Total", total))

	return formattedBill.String()
}

func (bill *bill) updateTip(tip float64) {
	bill.tip = tip
}

func (bill *bill) addItem(itemName string, price float64) {
	bill.items[itemName] = price
}

func (bill *bill) save() {
	data := []byte(bill.format())

	const billsFolder = "bills"

	_, readDirError := os.ReadDir(billsFolder)
	if readDirError != nil {
		createDirError := os.Mkdir(billsFolder, os.ModeDir)
		if createDirError != nil {
			panic(createDirError)
		}
	}

	for i := 0; i <= 100; i++ {
		fileName := fmt.Sprintf("%s/%s-%d.txt", billsFolder, strings.ToLower(bill.name), i)
		existingFile, _ := os.ReadFile(fileName)

		if existingFile == nil {
			createFileError := os.WriteFile(fileName, data, 0644)
			if createFileError != nil {
				panic(createFileError)
			}

			break
		}
	}

}
