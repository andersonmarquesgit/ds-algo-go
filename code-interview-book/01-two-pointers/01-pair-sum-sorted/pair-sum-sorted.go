package main

import "fmt"

// Time Complexity = O(nˆ2)
func findPairSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}

	return nil
}

// Time Complexity = O(n)
func findPairSumWithTwoPointers(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	nums := []int{-5, -2, 3, 4, 6}
	target := 7
	fmt.Println(findPairSum(nums, target))
	fmt.Println(findPairSumWithTwoPointers(nums, target))
}
