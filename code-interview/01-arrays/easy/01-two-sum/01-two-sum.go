package main

import "fmt"

// Question: https://leetcode.com/problems/two-sum/
/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
*/

// Step 1: Verify the constraints
// Are all numbers positive or can there be negative numbers? Positive [1, 3, 7, 9, 2]
// Are there duplicate numbers in the array? No, there are no duplicates
// Will there always be a solution available? No, there may not always be a solution. So in that case, we can return nil
// What do we return if there is no solution? Return nil
// Can there be multiple pairs that add up to the target? No, only one pair of numbers will add up to the target

// The first solution or brutal force is make the function verify each element on array
func findTwoSum(nums []int, target int) []int { // O(nˆ2)
	if len(nums) <= 1 {
		return nil
	}

	//nums=[1, 3, 7, 9, 2]  target=11
	for firstIdxNumber := 0; firstIdxNumber < len(nums); firstIdxNumber++ {
		numberToFind := target - nums[firstIdxNumber]
		for secondIdxNumber := firstIdxNumber + 1; secondIdxNumber < len(nums); secondIdxNumber++ {
			if numberToFind == nums[secondIdxNumber] {
				return []int{firstIdxNumber, secondIdxNumber}
			}
		}
	}

	return nil
}

// Optimizing function
func findTwoOptimizing(nums []int, target int) []int { // O(n)
	if len(nums) <= 1 {
		return nil
	}

	// nums=[1, 3, 7, 9, 2]  target=11
	// idxNumber = 0, number = 1, numberToFindMap[10] = 0
	// idxNumber = 1, number = 3, numberToFindMap[8] = 1
	// ..
	// idxNumber = 3, number = 9, numberToFindMap[2] = 3
	// idxNumber = 4, number = 2, yes we have number on Map, so return []int{numberToFindMap[number], idxNumber}
	numberToFindMap := make(map[int]int)
	for idxNumber, number := range nums {
		if _, ok := numberToFindMap[number]; ok {
			return []int{numberToFindMap[number], idxNumber}
		}

		numberToFind := target - number
		numberToFindMap[numberToFind] = idxNumber
	}

	return nil
}

func main() {
	nums := []int{1, 3, 7, 9, 2}
	target := 11

	fmt.Println(findTwoSum(nums, target))

	nums2 := []int{}
	fmt.Println(findTwoSum(nums2, target))

	nums3 := []int{1, 3, 7, 9, 2}

	fmt.Println(findTwoOptimizing(nums3, target))

	nums4 := []int{}
	fmt.Println(findTwoOptimizing(nums4, target))

}
