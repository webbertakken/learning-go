//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

package main

import "fmt"

//* Create a structure to store items and their security tag tag
//  - Security tags have two states: active (true) and inactive (false)

const (
	Active   = true
	Inactive = false
)

type SecurityTag = bool

type Product struct {
	name string
	tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activate(product *SecurityTag) {
	*product = Active
}

func deactivate(product *SecurityTag) {
	*product = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(products []Product) {
	for i := 0; i < len(products); i++ {
		deactivate(&products[i].tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	foo := Product{"Foo", Active}
	bar := Product{"Bar", Active}
	baz := Product{"Baz", Active}
	fuz := Product{"Fuz", Active}

	//  - Store them in a slice or array
	basket := []Product{foo, bar, baz, fuz}
	fmt.Println(basket)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&basket[1].tag)
	fmt.Println(basket)

	//  - Call the checkout() function to deactivate all tags
	checkout(basket)
	fmt.Println(basket)

}
