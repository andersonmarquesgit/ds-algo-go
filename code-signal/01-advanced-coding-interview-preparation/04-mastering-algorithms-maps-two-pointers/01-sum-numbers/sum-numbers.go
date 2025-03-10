package main

import "fmt"

func SumNumbers(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func SumNumbers2(n int) int {
	return n * (n + 1) / 2
}

func main() {
	result := SumNumbers(1000000)
	fmt.Println("Sum:", result)

	result2 := SumNumbers2(1000000)
	fmt.Println("Sum:", result2)
}
