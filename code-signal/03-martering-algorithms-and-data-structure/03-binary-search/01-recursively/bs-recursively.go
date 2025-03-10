package main

import "fmt"

func BinarySearch(arr []int, start, end, target int) int {
	if start > end {
		return -1 // Base case
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
	target := 3
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(arr, 0, len(arr)-1, target))

}
