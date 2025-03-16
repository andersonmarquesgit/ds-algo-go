package main

import (
	"fmt"
)

func quickSort(sizes []int, left, right int) {
	if left < right {
		pi := partition(sizes, left, right)
		quickSort(sizes, left, pi-1)  // Recursively sort left side
		quickSort(sizes, pi+1, right) // Recursively sort right side
	}
}

func partition(sizes []int, left, right int) int {
	pivot := sizes[right]
	i := left
	for j := left; j < right; j++ {
		if sizes[j] < pivot {
			sizes[i], sizes[j] = sizes[j], sizes[i]
			i++
		}
	}
	sizes[i], sizes[right] = sizes[right], sizes[i]
	return i
}

func main() {

	arr := []int{2, 5, 4, 3, 8, 97, 34, 22, 4, 86, 2, 1, 12, 52, 74}

	fmt.Println("Original List: ", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println("List After Quick Sort: ", arr)
}
