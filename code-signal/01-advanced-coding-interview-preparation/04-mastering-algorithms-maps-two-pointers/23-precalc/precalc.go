package main

import (
	"fmt"
)

func Solve(arr []int, queries [][2]int) []int {
	// TODO: Your solution
	n := len(arr)
	maxSum := make([][]int, n)

	for i := range maxSum {
		maxSum[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		curSum := 0
		curMax := arr[i]
		for j := i; j < n; j++ {
			curSum += arr[j]
			if curSum > curMax {
				curMax = curSum
			}
			maxSum[i][j] = curMax
		}

	}

	result := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		result[i] = maxSum[l][r]
	}

	return result
}

func main() {
	arr := []int{3, 1, -4, 2, -5, 1, 3, 6}
	queries := [][2]int{{0, 4}, {2, 6}, {2, 7}}

	result := Solve(arr, queries)
	fmt.Println(result)
}
