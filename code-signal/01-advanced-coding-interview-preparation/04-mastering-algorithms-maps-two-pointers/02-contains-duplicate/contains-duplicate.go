package main

import (
	"fmt"
	"sort"
)

func ContainsDuplicate(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

/*
Now, even including the time it takes to sort the array (usually O(n log n)), this updated function is more efficient.
The time complexity of the sorting operation using Go's standard sorting algorithm is O(n log n).
After sorting, we only make a single pass through the array — an O(n) operation. Combined, we have O(n log n) + O(n).
However, since O(n log n) grows faster than O(n), it is more significant than O(n) for large n, and thus the overall
time complexity is roughly O(n log n).
*/
func ContainsDuplicate2(arr []int) bool {
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return true
		}
	}
	return false
}

func ContainsDuplicate3(arr []int) bool {
	visited := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if _, exists := visited[arr[i]]; exists {
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	result := ContainsDuplicate(arr)
	fmt.Println("Contains Duplicate:", result)

	result2 := ContainsDuplicate2(arr)
	fmt.Println("Contains Duplicate:", result2)

	result3 := ContainsDuplicate3(arr)
	fmt.Println("Contains Duplicate:", result3)
}
