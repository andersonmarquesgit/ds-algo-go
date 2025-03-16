package main

import (
	"fmt"
)

func main() {
	vaults := []int{9, 8, 7, 3, 2, 10}
	fmt.Println(FindMax(vaults)) // Output: 10

	vaults2 := []int{5, 4, 3, 1, 6}
	fmt.Println(FindMax(vaults2)) // Output: 6

	vaults3 := []int{3, 2, 1, 10, 9}
	fmt.Println(FindMax(vaults3)) // Output: 10
}

func FindMax(vaults []int) int {
	left, right := 0, len(vaults)-1
	// Implement iterative binary search

	for left < right {
		mid := left + (right-left)/2

		if vaults[mid] == vaults[right] {
			return vaults[right]
		} else {
			if vaults[mid] < vaults[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return vaults[left]
}
      