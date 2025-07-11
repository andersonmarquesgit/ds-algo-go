package main

import "fmt"

// Time Complexity: O(nˆ2)
func largestContainer(heights []int) int {
	var maxContainer int
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			maxCurrent := min(heights[i], heights[j]) * (j - i)
			maxContainer = max(maxContainer, maxCurrent)
		}
	}

	return maxContainer
}

func largestContainerWithTwoPointers(heights []int) int {
	left, right := 0, len(heights)-1
	maxContainer := 0

	for left < right {
		maxCurrent := min(heights[left], heights[right]) * (right - left)
		maxContainer = max(maxContainer, maxCurrent)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxContainer
}

func main() {
	heights1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(largestContainer(heights1))
	fmt.Println(largestContainerWithTwoPointers(heights1))

	heights2 := []int{2, 7, 8, 3, 7, 6}
	fmt.Println(largestContainer(heights2))
	fmt.Println(largestContainerWithTwoPointers(heights2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
