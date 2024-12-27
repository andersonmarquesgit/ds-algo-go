package main

import "fmt"

var calculations = 0

// O(2^n) - Exponential time
func fibonacci(n int) int {
	calculations++
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

var calculationsWithDynamicProgramming = 0

var cache = map[int]int{}

func fibonacciWithDynamicProgramming(n int) int {
	calculationsWithDynamicProgramming++
	if n <= 1 {
		return n
	}

	if _, ok := cache[n]; !ok {
		cache[n] = fibonacciWithDynamicProgramming(n-1) + fibonacciWithDynamicProgramming(n-2)
	}

	return cache[n]

}

// []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...}
func main() {
	fmt.Println("-- Fibonacci without dynamic programming --")
	fmt.Println(fibonacci(30))
	fmt.Println("Calculations: ", calculations)

	fmt.Println("")
	fmt.Println("-- Fibonacci with dynamic programming --")
	fmt.Println(fibonacciWithDynamicProgramming(30))
	fmt.Println("Calculations with dynamic programming: ", calculationsWithDynamicProgramming)
}
