package main

import "fmt"

func factorialRecursion(n int) int {
	if n <= 1 { // Base case
		return n
	}

	return n * factorialRecursion(n-1) // Recursive case
}

func factorialInteractive(n int) int {
	if n == 0 {
		return 0
	}

	fact := 1
	for i := n; i > 1; i-- {
		fact *= i
	}

	return fact
}

func main() {
	fmt.Println(factorialRecursion(0))
	fmt.Println(factorialRecursion(1))
	fmt.Println(factorialRecursion(4))
	fmt.Println(factorialRecursion(8))

	fmt.Println(factorialInteractive(0))
	fmt.Println(factorialInteractive(1))
	fmt.Println(factorialInteractive(4))
	fmt.Println(factorialInteractive(8))
}
