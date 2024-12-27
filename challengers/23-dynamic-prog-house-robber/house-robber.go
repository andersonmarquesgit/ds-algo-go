/*
198. House Robber
Medium
Topics
Companies
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected
and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money
you can rob tonight without alerting the police.



Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/

package main

import "fmt"

/*
	 We have a array nums in disorder
	 We can try the solution using liner execution - for {house 1 + house 1 + 2 (not adjacent)} - Time Complexity is O(n)
		because we sum all elements not adjacent in array

	 We can use the dynamic programming with caching and recursion
*/

// [1,2,3,1]
func rob(nums []int) int { // O(n)
	cache := make(map[int]int)
	return robRecursive(nums, 0, cache)
}

func robRecursive(nums []int, i int, cache map[int]int) int { // O(n)
	if i >= len(nums) {
		return 0
	}

	if val, ok := cache[i]; ok {
		return val
	}

	robCurrentHouse := nums[i] + robRecursive(nums, i+2, cache)
	robNextHouse := robRecursive(nums, i+1, cache)

	cache[i] = max(robCurrentHouse, robNextHouse)
	return cache[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(rob([]int{1, 2, 3, 1}))    // 4 - Explain: Rob house 1 (money = 1) and then rob house 3 (money = 3).
	//fmt.Println(rob([]int{2, 7, 9, 3, 1})) // 12 - Explain: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
	fmt.Println(rob([]int{2, 1, 1, 2})) // 4 - Explain: Rob house 1 (money = 2) and then rob house 4 (money = 2). Total amount you can rob = 2 + 2 = 4.
	// Because we can't rob house 2 and 3 because they are adjacent. Choice house 1 and 4. because they are not adjacent and the total amount is 4.
}
