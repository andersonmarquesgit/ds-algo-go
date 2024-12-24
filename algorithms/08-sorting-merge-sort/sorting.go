package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) == 1 { // base case
		return array
	}

	mid := len(array) / 2
	left := array[:mid]
	right := array[mid:]
	fmt.Println("left:", left)
	fmt.Println("right:", right)

	return merge(mergeSort(left), mergeSort(right)) // recursive call
}

func merge(left, right []int) []int { // O(n)
	result := make([]int, len(left)+len(right))

	leftIndex := 0
	rightIndex := 0
	k := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result[k] = left[leftIndex]
			leftIndex++
		} else {
			result[k] = right[rightIndex]
			rightIndex++
		}
		k++
	}

	for leftIndex < len(left) {
		result[k] = left[leftIndex]
		leftIndex++
		k++
	}

	for rightIndex < len(right) {
		result[k] = right[rightIndex]
		rightIndex++
		k++
	}

	return result
}

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println("Final result: ", mergeSort(array)) // [1 2 3 4 5 6 7 8]
}
