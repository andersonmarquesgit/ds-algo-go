package main

import "fmt"

func binarySearch(arr []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func FindPositions(arr []int, target int) []int {
	if len(arr) == 0 {
		return []int{-1, -1}
	}

	firstPosition := binarySearch(arr, 0, len(arr)-1, target)
	if firstPosition == -1 {
		return []int{-1, -1}
	}

	start, end := firstPosition, firstPosition
	temp1, temp2 := 0, 0

	for start != -1 {
		temp1 = start
		start = binarySearch(arr, 0, start-1, target)
	}
	start = temp1

	for end != -1 {
		temp2 = end
		end = binarySearch(arr, end+1, len(arr)-1, target)
	}
	end = temp2

	return []int{start, end}
}

func main() {
	// Test case 1
	// arr = [1, 3, 3, 5, 5, 5, 5, 10], target = 5
	// Output: [3, 6]

	arr1 := []int{1, 3, 3, 5, 5, 5, 5, 10}
	fmt.Println(FindPositions(arr1, 5)) // Expected: [3, 6]

	// Test case 2
	// arr = [1, 5, 5, 5, 10], target = 5
	// Output: [1, 3]
	arr2 := []int{1, 5, 5, 5, 10}
	fmt.Println(FindPositions(arr2, 5)) // Expected: [1, 3]

}
