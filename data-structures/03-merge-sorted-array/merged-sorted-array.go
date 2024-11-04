package main

import "fmt"

// Given two arrays, merge them into a single sorted array
// mergedSortedArray([0,3,4,31], [4,6,30]) => [0,3,4,4,6,30,31]

func mergedSortedArray(arr1, arr2 []int) []int { // O(n^2)
	// Create a new array to store the merged array
	mergedArray := make([]int, len(arr1)+len(arr2))

	for i := 0; i < len(arr1); i++ { // O(n)
		mergedArray[i] = arr1[i]
	}

	for i := 0; i < len(arr2); i++ { // O(n)
		mergedArray[len(arr1)+i] = arr2[i]
	}

	// Sort the merged array
	for i := 0; i < len(mergedArray); i++ { // O(n^2)
		for j := i + 1; j < len(mergedArray); j++ {
			if mergedArray[i] > mergedArray[j] {
				mergedArray[i], mergedArray[j] = mergedArray[j], mergedArray[i]
			}
		}
	}

	return mergedArray
}

func mergedSortedArray2(arr1, arr2 []int) []int { // O(n)
	mergedArray := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0

	for i < len(arr1) && j < len(arr2) { // O(n)
		if arr1[i] < arr2[j] {
			mergedArray[k] = arr1[i]
			i++
		} else {
			mergedArray[k] = arr2[j]
			j++
		}
		k++
	}

	// Adicionar elementos restantes de arr1
	for i < len(arr1) { // O(n)
		mergedArray[k] = arr1[i]
		i++
		k++
	}

	// Adicionar elementos restantes de arr2
	for j < len(arr2) { // O(n)
		mergedArray[k] = arr2[j]
		j++
		k++
	}

	return mergedArray
}

func main() {
	arr1 := []int{0, 3, 4, 31}
	arr2 := []int{4, 6, 30}
	merged := mergedSortedArray(arr1, arr2)
	fmt.Println(merged)

	merged2 := mergedSortedArray2(arr1, arr2)
	fmt.Println(merged2)
}
