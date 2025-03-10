package main

import (
	"fmt"
)

// Recursive function to count the digits of a positive number
func CountDigits(num int) int {
	if num < 10 {
		return 1
	}

	return 1 + CountDigits(num/10) // Recursive operation to scan further digits
}

// Recursive function to calculate the sum of digits of a positive number
func CountStars(number int) int {
	if number == 0 { // Base case: if number is 0, stop recursion
		return 0
	}
	// TODO: Add the recursive logic to sum up the digits of 'number'
	return number%10 + CountStars(number/10)
}

func main() {
	fmt.Println(CountDigits(9876))
	fmt.Println(CountStars(9876))
}
