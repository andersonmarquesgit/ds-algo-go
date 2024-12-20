package main

import "fmt"

func bubbleSort(nums []int) []int {
	isSorted := false

	for !isSorted {
		swapNumbers := false
		for i := 0; i < len(nums); i++ {
			if i == len(nums)-1 {
				break
			}

			if nums[i] > nums[i+1] {
				temp := nums[i+1]
				nums[i+1] = nums[i]
				nums[i] = temp
				swapNumbers = true
			}
		}

		if !swapNumbers {
			isSorted = true
		}
	}

	return nums
}

func main() {
	nums := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(bubbleSort(nums)) // [1 2 3 4 5 6 7 8]

	nums2 := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(bubbleSort(nums2)) // [0 1 2 4 5 6 44 63 87 99 283]
}
