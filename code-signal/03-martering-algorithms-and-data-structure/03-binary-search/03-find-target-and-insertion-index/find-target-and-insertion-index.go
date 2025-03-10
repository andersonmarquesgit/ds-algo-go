package main

import (
	"fmt"
)

func BinarySearch(arr []int, start, end, target int) int {
	if start > end {
		return start // Base case
	}
	mid := start + (end-start)/2 // Find the midpoint
	if arr[mid] == target {
		return mid // Target found
	}
	if arr[mid] > target {
		return BinarySearch(arr, start, mid-1, target) // Search the left half
	}
	return BinarySearch(arr, mid+1, end, target) // Search the right half
}

func main() {
	numbers := []int{1, 3, 5, 6, 9, 12, 15, 18, 21}
	searchNumber := 14
	targetIndex := BinarySearch(numbers, 0, len(numbers)-1, searchNumber)
	if targetIndex != -1 {
		fmt.Printf("Number %d is at index: %d\n", searchNumber, targetIndex)
	} else {
		fmt.Printf("Number %d not found. But the insertion index is %d\n", searchNumber, targetIndex)
	}
}
