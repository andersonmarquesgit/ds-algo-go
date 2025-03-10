package main

import "fmt"

func findQuadSum(targetSum int, numbers []int) []int {
	sumMap := make(map[int][][2]int)

	// Step 1: Populate map with the sums of all pairs.
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			pairwiseSum := numbers[i] + numbers[j]
			sumMap[pairwiseSum] = append(sumMap[pairwiseSum], [2]int{i, j})
		}
	}

	// Step 2: Search for complement pairs.
	for sum, pairs1 := range sumMap {
		complement := targetSum - sum
		if pairs2, found := sumMap[complement]; found {
			for _, pair1 := range pairs1 {
				for _, pair2 := range pairs2 {
					a, b := pair1[0], pair1[1]
					c, d := pair2[0], pair2[1]

					// Ensure all indices are distinct.
					if a != c && a != d && b != c && b != d {
						return []int{numbers[a], numbers[b], numbers[c], numbers[d]}
					}
				}
			}
		}
	}

	return []int{}
}

func main() {
	numbers := []int{5, 15, 2, 7, 8, 4}
	target := 24
	result := findQuadSum(target, numbers)
	fmt.Println(result) // Output: [5 7 8 4]
}
