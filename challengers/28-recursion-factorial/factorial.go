package main

import "fmt"

func FirstFactorial(num int) int { // O(n)
	if num == 1 { // base case
		return 1
	}

	return num * FirstFactorial(num-1) // recursive case
}

func main() {
	fmt.Println(FirstFactorial(4))
}

// FirstFactorial(4)
// 4 * FirstFactorial(3)
// 4 * 3 * FirstFactorial(2)
// 4 * 3 * 2 * FirstFactorial(1)
// 4 * 3 * 2 * 1
