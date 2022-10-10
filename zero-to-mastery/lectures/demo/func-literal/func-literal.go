package main

import "fmt"

func add(a, b int) int {
  return a + b
}

func subtract(a, b int) int {
  return a - b
}

type OperatorFn func(a, b int) int

func compute(a, b int, operatorFn OperatorFn) int {
  fmt.Printf("Running a computation with %v and %v\n", a, b)
  return operatorFn(a, b)
}

func main() {
  fmt.Println("3 + 3 =", compute(3, 3, add))
  fmt.Println("3 - 3 =", compute(3, 3, subtract))
  fmt.Println("3 / 3 =", compute(3, 3, func(a, b int) int { return a / b }))

  mul := func(a, b int) int { return a * b }
  fmt.Println("3 * 3 =", compute(3, 3, mul))
}
