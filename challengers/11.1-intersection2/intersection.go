/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000


Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]int) // Map for the index int

	// Create a map to store the index of each int in nums1
	for _, v := range nums1 {
		set[v]++
	}

	fmt.Println(set)

	result := []int{}

	// Iterate over the map and append the int that exists in nums2
	for _, v := range nums2 {
		if set[v] > 0 {
			result = append(result, v)
			set[v]--
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2)) // [2, 2]

	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums3, nums4)) // [9, 4]
}
