package main

import "fmt"

/*
42. Trapping Rain Water
Hard
Topics
Companies
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it can trap after raining.



Example 1:


Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
In this case, 6 units of rain water (blue section) are being trapped.
Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9


Constraints:

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

/*
Time complexity: O(n^2 )
Space complexity: O(1)
*/
func getTrappedRainwater(heights []int) int {
	totalRainwater := 0

	for p := 0; p < len(heights); p++ {
		maxL := 0
		maxR := 0
		currentHeight := heights[p]

		for leftP := p - 1; leftP >= 0; leftP-- {
			if heights[leftP] > maxL {
				maxL = heights[leftP]
			}
		}

		for rightP := p + 1; rightP < len(heights); rightP++ {
			if heights[rightP] > maxR {
				maxR = heights[rightP]
			}
		}

		currentWater := min(maxL, maxR) - currentHeight

		if currentWater >= 0 {
			totalRainwater += currentWater
		}
	}

	return totalRainwater
}

/*
Time complexity: O(n)
Space complexity: O(1)
*/
func getTrappedRainwaterOptimized(heights []int) int {
	totalRainwater := 0
	leftP := 0
	rightP := len(heights) - 1
	maxLeft := 0
	maxRight := 0

	for leftP < rightP {
		currentWater := 0
		currentHeight := 0

		if heights[leftP] <= heights[rightP] {
			currentHeight = heights[leftP]
			if maxLeft > currentHeight {
				currentWater = maxLeft - currentHeight
				totalRainwater += currentWater
			} else {
				maxLeft = currentHeight
			}
			leftP++
		} else {
			currentHeight = heights[rightP]
			if maxRight > currentHeight {
				currentWater = maxRight - currentHeight
				totalRainwater += currentWater
			} else {
				maxRight = currentHeight
			}
			rightP--
		}
	}

	return totalRainwater
}

func min(p1, p2 int) int {
	if p1 < p2 {
		return p1
	}

	return p2
}

func main() {
	heights := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
	fmt.Println(getTrappedRainwater(heights))

	heights2 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(getTrappedRainwater(heights2))

	heights3 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(getTrappedRainwater(heights3))

	fmt.Println("---- Using func optimized ----")
	fmt.Println(getTrappedRainwaterOptimized(heights))
	fmt.Println(getTrappedRainwaterOptimized(heights2))
	fmt.Println(getTrappedRainwaterOptimized(heights3))
}
