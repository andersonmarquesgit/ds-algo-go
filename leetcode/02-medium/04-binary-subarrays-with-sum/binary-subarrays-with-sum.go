package main

import "fmt"

/*
930. Binary Subarrays With Sum
Solved
Medium
Topics
Companies
Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.

A subarray is a contiguous part of the array.



Example 1:

Input: nums = [1,0,1,0,1], goal = 2
Output: 4
Explanation: The 4 subarrays are bolded and underlined below:
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
Example 2:

Input: nums = [0,0,0,0,0], goal = 0
Output: 15


Constraints:

1 <= nums.length <= 3 * 104
nums[i] is either 0 or 1.
0 <= goal <= nums.length
*/

// O(n^2)
func numSubArraysWithSum(nums []int, goal int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		currentSum := 0

		for j := i; j < len(nums); j++ {
			currentSum += nums[j]

			if currentSum == goal {
				count++
			}
		}
	}

	return count
}

func numSubArraysWithSumOptimized(nums []int, goal int) int {
	count := 0
	prefixSum := 0
	freq := make(map[int]int)
	freq[0] = 1 // Inicializa para lidar com subarrays que começam do índice 0

	for _, num := range nums {
		prefixSum += num

		// Se (prefixSum - goal) já foi visto antes, significa que encontramos subarrays válidos
		if val, exists := freq[prefixSum-goal]; exists {
			count += val
		}

		// Atualiza o hash map com a nova soma prefixada
		freq[prefixSum]++
	}

	return count
}

func main() {
	nums := []int{1, 0, 1, 0, 1}
	goal := 2

	fmt.Println(numSubArraysWithSum(nums, goal))
	fmt.Println(numSubArraysWithSumOptimized(nums, goal))

	nums2 := []int{0, 0, 0, 0, 0}
	goal2 := 0

	fmt.Println(numSubArraysWithSum(nums2, goal2))
	fmt.Println(numSubArraysWithSumOptimized(nums2, goal2))
}
