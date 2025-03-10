package main

import (
	"fmt"
	"math"
)

func DiagonalTraverseAndSquares(matrix [][]int) []int {
	// TODO: we need the rows and cols size
	rows := len(matrix)
	cols := len(matrix[0])
	traverse := []int{}
	result := []int{}

	// TODO: we need the traverse slice for store
	// TODO: for the control direction we need dir = 1 (down-left) and dir = -1 (up-right)

	// For dir = 1 (down-left) we have this cases:
	// if row == rows - 1 (is limit for row) -> col += 1 and dir = -1 (IS VERY IMPORTANT CHECK FIRST THE BOUNDERING CASE, BECAUSE OF ERROR WITH INDEX OF OUT BOUND)
	// else if col == 0 -> row += 1 and dir = -1
	// else row += 1 and col -= 1

	// For dir = -1 (up-right) we this cases:
	// if col == cols - 1 (is limit for col) -> row += 1 and dir = 1  (IS VERY IMPORTANT CHECK FIRST THE BOUNDERING CASE, BECAUSE OF ERROR WITH INDEX OF OUT BOUND)
	// else if row == 0 -> col += 1 and dir = 1
	// else row -= 1 and col += 1

	// TODO: we need row and col variables
	row, col, dir := 0, 0, 1

	for i := 0; i < rows*cols; i++ {
		traverse = append(traverse, matrix[row][col])

		if dir == 1 {
			if row == rows-1 {
				col += 1
				dir = -1
			} else if col == 0 {
				row += 1
				dir = -1
			} else {
				row += 1
				col -= 1
			}
		} else {
			if col == cols-1 {
				row += 1
				dir = 1
			} else if row == 0 {
				col += 1
				dir = 1
			} else {
				row -= 1
				col += 1
			}
		}
	}

	// TODO: we need verify each element in traverse if are perfect squares and store the position 1-indexed on result
	for idx, number := range traverse {
		root := math.Sqrt(float64(number))
		if root == float64(int(root)) {
			result = append(result, idx+1)
		}
	}

	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	result := DiagonalTraverseAndSquares(matrix)
	fmt.Println(result) // Expected Output: [1, 6, 9]
}
