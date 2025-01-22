package main

import "fmt"

func quickSelect(nums []int, left, right, idxToFind int) int {
	if left < right {
		pivotPosition := right
		partitionIndex := partition(nums, left, right, pivotPosition)

		if partitionIndex == idxToFind {
			return nums[partitionIndex]
		} else if idxToFind < partitionIndex {
			return quickSelect(nums, left, partitionIndex-1, idxToFind)
		} else {
			return quickSelect(nums, partitionIndex+1, right, idxToFind)
		}
	}

	return -1
}

func partition(nums []int, left, right, pivotPosition int) int {
	pivotValue := nums[pivotPosition]
	partitionIndex := left

	for j := left; j < right; j++ {
		if nums[j] < pivotValue {
			swap(nums, j, partitionIndex)
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
	k := 4
	nums := []int{5, 3, 1, 6, 4, 2}
	fmt.Println(quickSelect(nums, 0, len(nums)-1, k))
}
