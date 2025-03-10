package main

import (
	"fmt"
	"math"
)

func SumNumbers(queries [][2]int) []int64 {
	// TODO: We need a loop in queries
	// TODO: We need find the min and max between a and b, for example: 1(a) and 5(b) => The max is 5 and min is 1
	// TODO: We can calculate the sum using optimization technique with min and max => max * (max + 1)/2 => 5 * (5 + 1)/2 => 30/2 => 15
	// and min => min * (min + 1)/2. So now the difference sumMax - sumMin is the result
	// TODO: We can store the result in a slice []int64 and return

	result := []int64{}

	for _, query := range queries {
		min := math.Min(float64(query[0]), float64(query[1]))
		max := math.Max(float64(query[0]), float64(query[1]))
		sumPointsMin := min * (min - 1) / 2
		sumPointsMax := max * (max + 1) / 2

		result = append(result, int64(sumPointsMax-sumPointsMin))
	}

	return result
}

func main() {
	queries := [][2]int{
		{2, 5},
		{5, 1},
		{1, 1},
		{500, 500},
		{1, 500},
		{123, 321},
	}

	results := SumNumbers(queries)
	for _, result := range results {
		fmt.Println(result)
	}
}
