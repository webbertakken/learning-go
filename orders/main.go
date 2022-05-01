package main

import "fmt"

func main() {
	myBill := newBill("first order bill")

	myBill.updateTip(10)
	myBill.addItem("pancakes", 12)

	fmt.Println(myBill.format())
}
