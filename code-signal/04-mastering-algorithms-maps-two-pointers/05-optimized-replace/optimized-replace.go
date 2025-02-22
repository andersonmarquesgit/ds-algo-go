package main

import (
	"fmt"
	"sort"
)

// Helper struct to store original indices of B
type Pair struct {
	value int
	index int
}

func optimizedReplace(A []int, B []int) []int {
	n := len(A)
	result := make([]int, n)

	// Step 1: Store original indices and sort B
	sortedB := make([]Pair, n)
	for i := 0; i < n; i++ {
		sortedB[i] = Pair{value: B[i], index: i}
	}
	sort.Slice(sortedB, func(i, j int) bool {
		return sortedB[i].value < sortedB[j].value
	})

	// Step 2: Find the closest neighbor for each element in sorted B
	closestMap := make(map[int]int) // Maps original index to closest index

	for i := 0; i < n; i++ {
		currIndex := sortedB[i].index
		closestIndex := -1
		minDiff := int(1e9) + 1 // Large initial value

		// Check left neighbor
		if i > 0 {
			leftIndex := sortedB[i-1].index
			diff := abs(sortedB[i].value - sortedB[i-1].value)
			if diff < minDiff {
				minDiff = diff
				closestIndex = leftIndex
			}
		}

		// Check right neighbor
		if i < n-1 {
			rightIndex := sortedB[i+1].index
			diff := abs(sortedB[i].value - sortedB[i+1].value)
			if diff < minDiff {
				minDiff = diff
				closestIndex = rightIndex
			}
		}

		// Store closest mapping
		closestMap[currIndex] = closestIndex
	}

	// Step 3: Construct result array based on closestMap
	for i := 0; i < n; i++ {
		result[i] = A[closestMap[i]]
	}

	return result
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Main function to test the solution
func main() {
	A := []int{10, 20, 30, 40, 50}
	B := []int{7, 5, 1, 2, 4}
	result := optimizedReplace(A, B)
	fmt.Println(result) // Expected Output: [20, 50, 40, 30, 20]
}
