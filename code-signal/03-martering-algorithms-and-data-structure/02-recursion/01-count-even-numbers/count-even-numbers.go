package main

import (
	"fmt"
)

// Function to count the number of digits in a number using recursion
func countDigits(num int) int {
	if num == 0 {
		return 0
	} else {
		// Recursive case: reduce the number by dividing by 10 and count the rest
		// TODO: Count only the even numbers
		lastDigit := num % 10
		if lastDigit%2 == 0 {
			return 1 + countDigits(num/10)
		}
		return countDigits(num / 10) // Ignorar dígitos ímpares
	}
}

func main() {
	fmt.Println("The number of digits in the 'diameter of Mars':", countDigits(7623))
}
