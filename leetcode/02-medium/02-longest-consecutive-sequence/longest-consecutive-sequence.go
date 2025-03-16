package main

import (
	"fmt"
	"sort"
)

/*
128. Longest Consecutive Sequence
Solved
Medium
Topics
Companies
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
Example 3:

Input: nums = [1,0,1,2]
Output: 3

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

// O(n log n)
func LongestConsecutiveSequence(nums []int) int {
	longestSequence := 1
	currentSequence := 1

	if len(nums) == 0 {
		return 0
	}

	// O(n log n)
	sort.Ints(nums)

	// O(n)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				currentSequence++
			} else {
				currentSequence = 1
			}

			longestSequence = max(longestSequence, currentSequence)
		}
	}

	return longestSequence
}

// O(n)
func LongestConsecutiveSequenceOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	// O(n)
	for _, num := range nums {
		numSet[num] = true
	}

	longestSequence := 0

	// O(n)
	for num := range numSet {
		// Só iniciamos a contagem se for o início de uma sequência
		if !numSet[num-1] {
			currentNum := num
			currentSequence := 1

			// Conta a sequência crescente
			for numSet[currentNum+1] {
				currentNum++
				currentSequence++
			}

			longestSequence = max(longestSequence, currentSequence)
		}
	}

	return longestSequence
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(LongestConsecutiveSequence(nums1))
	fmt.Println(LongestConsecutiveSequenceOptimized(nums1))

	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(LongestConsecutiveSequence(nums2))
	fmt.Println(LongestConsecutiveSequenceOptimized(nums2))

	nums3 := []int{1, 0, 1, 2}
	fmt.Println(LongestConsecutiveSequence(nums3))
	fmt.Println(LongestConsecutiveSequenceOptimized(nums3))
}
