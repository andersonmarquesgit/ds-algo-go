package main

import (
	"fmt"
)

// MinRemovalsToBalance creates a function that determines the minimum number of bracket removals needed for a valid string.
func MinRemovalsToBalance(brackets string) int {
	// TODO: Initialize an empty stack to act as the stack
	stack := []rune{}

	// TODO: Initialize a variable to count removals
	countRemoval := 0

	// TODO: Iterate through each bracket in the input string
	for _, char := range brackets {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' { // is closing ')'
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Remove '(' correspondent
			} else {
				countRemoval++ // `)` without correspondent
			}
		}
	}

	// Adicionamos os `(` restantes na pilha ao contador de remoções
	countRemoval += len(stack)

	// TODO: Add conditions to handle the opening and closing brackets appropriately using stack operations

	// TODO: Return the count of brackets that need to be removed to make the string valid
	return countRemoval // Placeholder return
}

// Example usage
func main() {
	invalidParentheses := "()))(()"
	removalsNeeded := MinRemovalsToBalance(invalidParentheses)
	fmt.Println(removalsNeeded) // Expected output: 3
}
