package main

import (
	"fmt"
	"sort"
)

func FindPairs(numbers []int, target int) [][]int { // Using two pointer technique
	sort.Ints(numbers)
	// [-2, 1, 2, 3, 5, 8]
	left := 0
	right := len(numbers) - 1
	pairs := [][]int{}

	// Using two-pointer technique to find pairs
	for left < right {
		total := numbers[left] + numbers[right]

		if total == target {
			// Add the pair to the result slice if the sum matches the target
			pairs = append(pairs, []int{numbers[left], numbers[right]})
			left++  // Move left pointer one step to the right
			right-- // Move right pointer one step to the left
		} else if total < target {
			left++ // Move left pointer one step to the right to increase the sum
		} else {
			right-- // Move right pointer one step to the left to decrease the sum
		}
	}
	return pairs // Return the slice of pairs that sum up to the target
}

func FindPairs2(numbers []int, target int) [][]int { // Using complement technique
	complementMap := make(map[int]int)
	result := [][]int{}

	for _, num := range numbers {
		complement := target - num
		if _, exists := complementMap[num]; exists {
			result = append(result, []int{complement, num})
		} else {
			complementMap[complement] = num
		}
	}

	sort.Slice(result, func(i, j int) bool {
		numI := result[i][0]
		numJ := result[j][0]

		return numI < numJ
	})

	return result
}

// Example usage
func main() {
	numbers := []int{1, 3, 5, 2, 8, -2}
	result := FindPairs(numbers, 6)
	for _, pair := range result {
		fmt.Printf("[%d, %d]\n", pair[0], pair[1])
	}
	// Output: [-2, 8]
	//         [1, 5]

	result2 := FindPairs2(numbers, 6)
	for _, pair := range result2 {
		fmt.Printf("[%d, %d]\n", pair[0], pair[1])
	}
}
