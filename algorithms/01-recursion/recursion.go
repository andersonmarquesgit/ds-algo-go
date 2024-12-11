/*
Given a number N return the index value of the Fibonacci sequence, where the sequence is:
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
The pattern of the sequence is that each value is the sum of the 2 previous values, that means that for N=5 → 2+3
*/
package main

import "fmt"

func fibonacciInteractive(n int) int {
	if n < 2 { // Base case
		return n
	}

	fib := []int{0, 1}

	for i := 2; i <= n; i++ { // O(n) - Linear time - Why? Because we are iterating over the n elements
		fib = append(fib, fib[i-1]+fib[i-2]) // Append the sum of the last two elements
	}

	fmt.Println(fib)

	return fib[n]
}

func fibonacciRecursive(n int) int {
	if n < 2 { // Base case
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2) // O(2^n) - Exponential time - Why? Because we are calling the function twice

}

func main() {
	fmt.Println("-- Fibonacci interactive --")
	fmt.Println(fibonacciInteractive(2))
	fmt.Println(fibonacciInteractive(3))

	fmt.Println("-- Fibonacci recursive --")
	fmt.Println(fibonacciRecursive(2))
	fmt.Println(fibonacciRecursive(3))
}
