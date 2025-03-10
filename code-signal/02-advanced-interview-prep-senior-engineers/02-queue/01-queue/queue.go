package main

import (
	"fmt"
)

func main() {
	var queue []string

	// Enqueue items
	queue = append(queue, "Apple")
	queue = append(queue, "Banana")
	queue = append(queue, "Cherry")

	// Dequeue an item
	if len(queue) > 0 {
		firstElement := queue[0]
		fmt.Println(firstElement) // Expects "Apple"
		queue = queue[1:]
	}
}
