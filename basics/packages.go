package main

import (
	"fmt"
	"sort"
	"strings"
)

func packages() {
	// "strings" package
	greeting := "Hello there friends!"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "friends", "peeps"))

	fmt.Printf("Original value of greeting is still %q\n", greeting)
	fmt.Println("index of \"ll\" is", strings.Index(greeting, "ll")) // 2
	fmt.Println(strings.Split(greeting, " "))                        // slice of 3 elements

	// "sort" package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages) // modifies the original slice
	fmt.Println("(sorted)", ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println("index of age 30 (after sorting) is", index)

	names := []string{"Yoshi", "Mario", "Peach", "Bowser", "Luigi"}
	sort.Strings(names) // modifies the original slice again
	fmt.Println(names)
	fmt.Println("Index of \"Bowser\" is", sort.SearchStrings(names, "Bowser"))
}
