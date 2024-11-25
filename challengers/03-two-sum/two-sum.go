package main

import "fmt"

/*
*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/
func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i, v := range nums {
		if idx, ok := set[v]; ok {
			return []int{idx, i}
		}
		set[target-v] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

	nums2 := []int{20, 3, 11, 1, 4, 9, 99}
	target2 := 10
	fmt.Println(twoSum(nums2, target2))
}
