package main

import "fmt"

func BinarySearchIterative(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2 // Calculate the midpoint to divide the search area

		if arr[mid] == target {
			return mid // Target found, return the index
		}

		if arr[mid] < target {
			start = mid + 1 // Move start to right of mid to search the right half
		} else {
			end = mid - 1 // Move end to left of mid to search the left half
		}
	}
	return -1 // Target not found, return -1
}

func main() {
	target := 3
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearchIterative(arr, target))

}
