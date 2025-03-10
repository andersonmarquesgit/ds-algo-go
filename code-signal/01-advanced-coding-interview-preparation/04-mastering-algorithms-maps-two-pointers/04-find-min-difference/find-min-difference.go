package main

import (
	"fmt"
	"math"
	"sort"
)

func FindMinDifference(nums []int) int {
	// TODO: implement the function to find minimum absolute difference among all pairs.
	sort.Ints(nums)

	smaller := math.MaxInt

	for i := 1; i < len(nums); i++ {
		currentSmaller := nums[i] - nums[i-1]
		if currentSmaller < smaller {
			smaller = currentSmaller
		}
	}
	return smaller
}

func main() {
	nums := []int{3, 8, 15, 10}
	fmt.Println(FindMinDifference(nums)) // Expected output: 2
}
