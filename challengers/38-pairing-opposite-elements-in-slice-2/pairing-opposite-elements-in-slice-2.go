package main

import "fmt"

func Solution(numbers []int) [][]int {
	// TODO: implement the function to find all pairs of integers in the list
	// where each integer has its reverse counterpart also present in the list.
	result := [][]int{}
	seen := make(map[int]bool)

	for _, num := range numbers {
		seen[num] = true
	}

	for _, num := range numbers {
		reversed := reverse(num)

		if seen[reversed] {
			result = append(result, []int{num, reversed})
		}
	}

	return result
}

func reverse(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func main() {
	numbers := []int{12, 21, 34, 43, 56, 65, 78}
	fmt.Println(Solution(numbers))
}
