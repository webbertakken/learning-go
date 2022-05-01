package main

import (
	"fmt"
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
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
}

func (bill *bill) format() string {
	total := 0.00
	formattedBill := strings.Builder{}

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
