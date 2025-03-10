package main

/*
 Now, put your knowledge to the test by evaluating an expression in Reverse Polish notation
(also known as postfix notation), where operands precede the operator. In other words,
the operator comes after the two numbers on which it is supposed to operate.

Your task is to write a function that processes such an expression and returns the result. Remember,
in a Reverse Polish Notation expression like "1 2 + 4 -", the operation is performed from left to right.
This means you first add 1 and 2, and then subtract 3 from the result, getting -1 as a final result.

Assume the expression consists solely of + and - operators, with spaces acting as separators between the elements.
*/
import (
	"fmt"
	"strconv"
	"strings"
)

func EvaluateReversePolishNotation(expression string) int {
	// TODO: Initialize a slice to simulate a stack for holding integer values
	stack := []int{}

	// TODO: Split the expression into tokens using whitespace as the delimiter
	expressionSplited := strings.Fields(expression)

	total := 0

	// TODO: Iterate over each token in the split expression
	for _, token := range expressionSplited {
		// TODO: If the token is an operator ('+' or '-'), pop the top two elements from the stack,
		// perform the corresponding operation, and push the result back onto the stack
		if token == "+" || token == "-" {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			// TODO: If the token is an operand, parse it to an integer and push it onto the stack
			if token == "+" {
				total = num1 + num2
			} else if token == "-" {
				total = num1 - num2
			}
			stack = append(stack, total)
		} else { // Is a number
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	// TODO: Return the final result, which should be the only element left in the stack
	return total // Placeholder return statement
}

func main() {
	// The expression "1 2 + 4 -" is "(1 + 2) - 4"
	fmt.Println(EvaluateReversePolishNotation("1 2 + 4 -")) // Expected output: -1
}
