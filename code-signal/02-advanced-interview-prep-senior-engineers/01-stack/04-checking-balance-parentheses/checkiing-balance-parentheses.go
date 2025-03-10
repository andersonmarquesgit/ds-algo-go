package main

import (
	"fmt"
)

func IsValidExpression(expression string) bool {
	stack := []rune{}
	matchingParen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range expression {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if closing, found := matchingParen[char]; found {
			// TODO: Determine if the stack is empty OR the last character does not match the corresponding opening character
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	// TODO: Check if the stack is empty, indicating that the expression is balanced
	return len(stack) == 0 // Modify this line appropriately
}

func main() {
	// Example usage
	sampleExpression := "([a+b]{c+d})"
	fmt.Println(IsValidExpression(sampleExpression)) // Expected output: true

	badExpression := "([a+b]{c+d}))"
	fmt.Println(IsValidExpression(badExpression)) // Expected output: false
}
