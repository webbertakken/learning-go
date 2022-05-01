package main

import (
	"fmt"
)

func pointers() {
	name := "foo"
	namePointer := &name
	fmt.Println("The memory address:", namePointer)         // 0xc00004a6b0
	fmt.Println("The value at that address:", *namePointer) // bar

	updateName(namePointer)
	fmt.Println("The updated value at that address:", name) // baz
}

func updateName(pointer *string) {
	*pointer = "bar"
}
