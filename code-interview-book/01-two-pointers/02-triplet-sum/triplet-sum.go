package main

import (
	"fmt"
	"sort"
)

// Time Complexity: O(nˆ3)
func findTripletSum(arr []int, target int) [][]int {
	n := len(arr)
	var triplets [][]int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == target {
					triplet := []int{arr[i], arr[j], arr[k]}
					sort.Ints(triplet)
					triplets = append(triplets, triplet)
				}
			}
		}
	}

	return triplets
}

func findTripletSumTwWithTwoPointers(arr []int, target int) [][]int {
	n := len(arr)
	var triplets [][]int
	sort.Ints(arr)

	for i := 0; i < n; i++ {
		// Test for 'a's duplicates
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		// Find all pairs that sum to a target 'a'
		pairs := findPairSum(arr[i+1:], target-arr[i])
		for _, pair := range pairs {
			triplet := []int{arr[i], pair[0], pair[1]}
			sort.Ints(triplet)
			triplets = append(triplets, triplet)
		}
	}

	return triplets
}

func findPairSum(arr []int, target int) [][]int {
	n := len(arr)
	var pairs [][]int

	left, right := 0, n-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			pairs = append(pairs, []int{arr[left], arr[right]})
			left++
			for left < right && arr[left] == arr[left-1] {
				left++
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return pairs
}

func main() {
	target := 0
	arr := []int{0, -1, 1, -3, 2}

	fmt.Println(findTripletSum(arr, target))
	fmt.Println(findTripletSumTwWithTwoPointers(arr, target))
}
