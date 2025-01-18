package main

import (
	"fmt"
)

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints
of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
In this case, the max area of water (blue section) the container can contain is 49.

Because the vertical lines used are 8 [index 1] and 7 [index 8], when area calculated is 7x7 = 49


Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

*/

/*
Time complexity: O(n^2)
Space complexity: O(1)
*/
func getMaxWaterContainer(heights []int) int { // O(n^2)
	maxArea := 0
	area := 0

	for p1 := 0; p1 < len(heights); p1++ {
		for p2 := p1 + 1; p2 < len(heights); p2++ {
			left := heights[p1]
			right := heights[p2]
			height := min(left, right)
			width := p2 - p1
			area = height * width

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

/*
Time complexity: O(n)
Space complexity: O(n)
*/
func getMaxWaterContainerOptimized(heights []int) int {
	maxArea := 0
	p1 := 0
	p2 := len(heights) - 1

	for p1 < p2 {
		height := min(heights[p1], heights[p2])
		width := p2 - p1
		area := height * width

		if area > maxArea {
			maxArea = area
		}

		if heights[p1] <= heights[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
}

func min(left, right int) int {
	if left < right {
		return left
	}

	return right
}

func main() {
	heights := []int{7, 1, 2, 3, 9}
	fmt.Println(getMaxWaterContainer(heights))

	heights2 := []int{}
	fmt.Println(getMaxWaterContainer(heights2))

	heights3 := []int{1}
	fmt.Println(getMaxWaterContainer(heights3))

	fmt.Println("---Using func optimized ----")
	fmt.Println(getMaxWaterContainerOptimized(heights))
	fmt.Println(getMaxWaterContainerOptimized(heights2))
	fmt.Println(getMaxWaterContainerOptimized(heights3))

	heights4 := []int{4, 8, 1, 2, 3, 9}
	fmt.Println(getMaxWaterContainerOptimized(heights4))
}
