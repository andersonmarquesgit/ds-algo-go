package main

import "fmt"

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
