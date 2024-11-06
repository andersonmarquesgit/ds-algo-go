package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums) // [5 6 7 1 2 3 4]

	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	fmt.Println(nums2) // [3 99 -1 -100]
}
