package main

import "fmt"

/*
Move Zeroes

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Could you minimize the total number of operations done?
*/

func moverZeroes(nums []int) {
	// Create a variable to store the index of the last non-zero element
	lastNonZeroFoundAt := 0

	//[0,1,0,3,12]
	// i = 1 | num = 1 troca | lastNonZeroFoundAt = 1 e [1,0,0,3,12]
	// i = 2 | num = 0 | lastNonZeroFoundAt = 1
	// i = 3 | num = 3 troca | lastNonZeroFoundAt = 2 e [1,3,0,0,12]
	// i = 4 | num = 12 troca | lastNonZeroFoundAt = 3 e [1,3,12,0,0]
	// Iterate over the array
	for i := 0; i < len(nums); i++ {
		// If the current element is not zero
		if nums[i] != 0 {
			// Swap the current element with the last non-zero element
			nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
			// Increment the index of the last non-zero element
			lastNonZeroFoundAt++
		}
	}

	fmt.Println(nums)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moverZeroes(nums)

	nums = []int{0}
	moverZeroes(nums)
}
