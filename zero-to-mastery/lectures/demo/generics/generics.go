package main

import "fmt"

const (
  Low = iota
  Medium
  High
)

type PriorityQueue[T comparable, V any] struct {
  items      map[T][]V
  priorities []T
}

func (pq *PriorityQueue[T, V]) Add(priority T, value V) {
  pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueue[T, V]) Next() (V, bool) {
  for i := 0; i < len(pq.priorities); i++ {
    priority := pq.priorities[i]
    items := pq.items[priority]
    if len(items) > 0 {
      next := items[0]
      pq.items[priority] = items[1:]
      return next, true
    }
  }
  var empty V
  return empty, false
}

func NewPriorityQueue[T comparable, V any](priorities []T) PriorityQueue[T, V] {
  return PriorityQueue[T, V]{items: make(map[T][]V), priorities: priorities}
}

func main() {
  queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})
  queue.Add(High, "H-1")
  queue.Add(Low, "L-1")

  fmt.Println(queue.Next())

  queue.Add(Medium, "M-1")
  queue.Add(Low, "L-2")
  queue.Add(High, "H-2")
  queue.Add(High, "H-3")
  queue.Add(Low, "L-3")
  queue.Add(Medium, "M-2")

  fmt.Println(queue.Next())
  fmt.Println(queue.Next())
  fmt.Println(queue.Next())
  fmt.Println(queue.Next())
  fmt.Println(queue.Next())
  fmt.Println(queue.Next())
  fmt.Println(queue.Next())
  fmt.Println(queue.Next())
}
