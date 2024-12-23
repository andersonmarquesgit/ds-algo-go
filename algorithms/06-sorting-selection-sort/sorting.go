package main

import "fmt"

func selectionSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		minorTemp := nums[i]
		swapPosition := i
		for j := i; j < length; j++ {
			if nums[j] < minorTemp {
				minorTemp = nums[j]
				swapPosition = j
			}
		}
		nums[swapPosition] = nums[i]
		nums[i] = minorTemp

	}

	return nums
}

func main() {
	nums := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	fmt.Println(selectionSort(nums))
}
