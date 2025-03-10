package main

import (
	"fmt"
)

func GetLongestSubarray(array []int, k int) []int {
	// TODO: We need the window slice, leftWindow and rightWindow pointer
	leftWindow, rightWindow := 0, 0
	longestSubarray := []int{}
	currentSum := 0

	// TODO: We need a loop in array using right
	for rightWindow < len(array) {
		currentSum += array[rightWindow]

		// If the sum is greater than `k`, we move `left`
		for currentSum > k && leftWindow <= rightWindow {
			currentSum -= array[leftWindow]
			leftWindow++
		}

		// If we find a sum equal `k`, verify if is the greater of subarray at now
		sizeWindow := rightWindow - leftWindow + 1
		if currentSum == k && sizeWindow > len(longestSubarray) {
			longestSubarray = append([]int{}, array[leftWindow:rightWindow+1]...) // Criar cópia
		}

		// We increase the window
		rightWindow++

	}

	return longestSubarray
}

/*
Test case - {1, 2, 3, 4, 5} | k=5
left=1
right=2
currentSum=5
array[rightWindow]=3
array[leftWindow]=1
longestSubarray=[2, 3]


Test case - {1, 2, 3, 4, 5} | k=12
left=2
right=5
currentSum=12
array[rightWindow]=5
array[leftWindow]=2
longestSubarray=[3, 4, 5]

*/

func main() {
	array := []int{1, 2, 3, 4, 5}
	k := 5
	result := GetLongestSubarray(array, k)
	fmt.Println(result) // Expected output: [2, 3]
}
