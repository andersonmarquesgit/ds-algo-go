package main

import "fmt"

// Time Complexity: O(nˆ2)
func pairSumUnsorted(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// Time Complexity: O(n)
func pairSumUnsortedOptimized(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok && i != j {
			return []int{i, j}
		}
	}

	return []int{}
}

// Time Complexity: O(n)
func pairSumUnsortedOptimized2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}

	return []int{}
}

func main() {
	nums := []int{-1, 3, 4, 2}
	target := 3
	fmt.Println(pairSumUnsorted(nums, target))
	fmt.Println(pairSumUnsortedOptimized(nums, target))
	fmt.Println(pairSumUnsortedOptimized2(nums, target))
}
