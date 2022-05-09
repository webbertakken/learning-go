package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	fmt.Println("Need cereal?")
	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yup", cereal, "boxes")
	}

	totalItems := 0
	for _, value := range shoppingList {
		totalItems += value
	}
	fmt.Println("Total items on the list", totalItems)
}
