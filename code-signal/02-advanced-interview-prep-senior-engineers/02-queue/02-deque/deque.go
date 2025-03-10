package main

import (
	"fmt"
)

func main() {
	var deque []string

	// Add items to both ends
	deque = append([]string{"Left end"}, deque...)
	deque = append(deque, "Middle")
	deque = append(deque, "Right end")

	// Remove an item from the right end
	if len(deque) > 0 {
		lastElement := deque[len(deque)-1]
		fmt.Println(lastElement) // Expects "Right end"
		deque = deque[:len(deque)-1]
	}

	// Remove an item from the left end
	if len(deque) > 0 {
		firstElement := deque[0]
		fmt.Println(firstElement) // Expects "Left end"
		deque = deque[1:]
	}
}
