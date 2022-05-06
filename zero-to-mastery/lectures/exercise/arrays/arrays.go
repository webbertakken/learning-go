//--Summary:

package main

import "fmt"

//  Create a program that can store a shopping list and print
//  out information about the list.
//

type Product struct {
	Name  string
	Price float64
}

type ShoppingList struct {
	Products [4]Product
}

func (l *ShoppingList) showInfo() {
	//* Print to the terminal:
	//  - The lastProduct item on the list
	lastProduct, _ := l.lastProduct()
	fmt.Println("The lastProduct product is: ", lastProduct.Name)

	//  - The total number of items
	numberOfProducts, _ := l.numberOfProducts()
	fmt.Println("Total number of products is: ", numberOfProducts)

	//  - The total cost of the items
	total, _ := l.totalCosts()
	fmt.Println("total", total)
}

func (l *ShoppingList) totalCosts() (float64, error) {
	total := 0.00

	for i := 0; i < len(l.Products); i++ {
		total += l.Products[i].Price
	}

	return total, nil
}

func (l *ShoppingList) numberOfProducts() (int, error) {
	products := 0

	for i := 0; i < len(l.Products); i++ {
		if l.Products[i].Name != "" {
			products += 1
		}
	}

	return products, nil
}

func (l *ShoppingList) lastProduct() (Product, error) {
	lastProduct := Product{}

	for i := 0; i < len(l.Products); i++ {
		if l.Products[i].Name != "" {
			lastProduct = l.Products[i]
		}
	}

	return lastProduct, nil
}

func main() {
	//--Requirements:
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//  - Products must include the price and the name
	shoppingList := ShoppingList{Products: [4]Product{}}

	//* Insert 3 products into the array
	shoppingList.Products[0] = Product{Name: "Ice tea", Price: 1.99}
	shoppingList.Products[1] = Product{Name: "Kiwis", Price: 2.99}
	shoppingList.Products[2] = Product{Name: "Expensive apples", Price: 3.99}
	shoppingList.showInfo()

	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList.Products[3] = Product{Name: "Linen bag", Price: 2.50}
	shoppingList.showInfo()
}
