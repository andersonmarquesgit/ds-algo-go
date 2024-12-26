package main

import "fmt"

func quicksort(nums []int, left, right int) []int {
	if left < right {
		pivot := right
		partitionIndex := partition(nums, left, right, pivot)
		quicksort(nums, left, partitionIndex-1)
		quicksort(nums, partitionIndex+1, right)
	}

	return nums
}

func partition(nums []int, left, right, pivot int) int {
	pivotValue := nums[pivot]
	partitionIndex := left

	for i := left; i < right; i++ {
		if nums[i] < pivotValue {
			swap(nums, i, partitionIndex)
			partitionIndex++
		}
	}

	swap(nums, right, partitionIndex)

	return partitionIndex
}

func swap(nums []int, firstIndex, secondIndex int) {
	temp := nums[firstIndex]
	nums[firstIndex] = nums[secondIndex]
	nums[secondIndex] = temp
}

func main() {
	nums := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(quicksort(nums, 0, len(nums)-1))

	nums2 := []int{3, 7, 8, 5, 2, 1, 9, 5, 4}
	fmt.Println(quicksort(nums2, 0, len(nums2)-1))
}
