package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool) // Map for the index int

	// Create a map to store the index of each int in nums1
	for _, n1 := range nums1 {
		set[n1] = true
	}

	result := []int{}

	// Iterate over the map and append the int that exists in nums2
	for _, n2 := range nums2 {
		if _, ok := set[n2]; ok {
			result = append(result, n2)
			delete(set, n2)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2)) // [2]

	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums3, nums4)) // [9, 4]
}
