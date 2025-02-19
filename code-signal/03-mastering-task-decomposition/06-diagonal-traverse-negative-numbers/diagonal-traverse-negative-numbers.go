package main

import "fmt"

func solution(matrix [][]int) [][]int {
	// TODO: First we need the rows and cols size
	// TODO: We need a map of slices int to store the position for negative number in matrix during first loop, this is help us to avoid a O(n^2) in matrix after give the traverse
	// TODO: We need a slice to store the traverse: {1, -2, 5, -9, -6, 3, -4, 7, 10, -11, 8, 12}
	// TODO: We can use two directions, dir = 1 (down-left) and dir = -1 (up-right)
	// TODO: Consider IF dir == 1 and col == 0 -> row += 1 and dir = -1
	// TODO: Consider IF dir == 1 and row == rows - 1 -> row += 1 and dir = -1
	// TODO: Consider IF dir == 1 and continue -> row += 1 and col -= 1
	// TODO: Consider IF dir == -1 and row == 0 -> col += 1 and dir = 1
	// TODO: Consider IF dir == -1 and col == cols-1 -> row += 1 and dir = 1
	// TODO: Consider IF dir == -1 and continue -> col += 1 and row -= 1

	rows := len(matrix)
	cols := len(matrix[0])
	positionsMap := make(map[int][]int)
	traverse := []int{}
	row, col, dir := 0, 0, -1

	for i := 0; i < rows*cols; i++ {
		currentNumber := matrix[row][col]
		traverse = append(traverse, currentNumber)

		if currentNumber < 0 {
			positionsMap[currentNumber] = []int{row + 1, col + 1}
		}

		if dir == 1 { // Moving down-left
			if row == rows-1 { // Hit bottom boundary
				col += 1
				dir = -1 // Switch direction to up-right
			} else if col == 0 { // Hit left boundary
				row += 1
				dir = -1 // Switch direction to up-right
			} else {
				row += 1
				col -= 1 // Continue down-left
			}
		} else {               // Moving up-right
			if col == cols-1 { // Hit right boundary
				row += 1
				dir = 1 // Switch direction to down-left
			} else if row == 0 { // Hit top boundary
				col += 1
				dir = 1 // Switch direction to down-left
			} else {
				row -= 1
				col += 1 // Continue up-right
			}
		}
	}

	result := [][]int{}

	for _, number := range traverse {
		if pos, exists := positionsMap[number]; exists {
			result = append(result, pos)
		}
	}

	return result
}

func main() {
	exampleMatrix := [][]int{
		{1, -2, 3, -4},
		{5, -6, 7, 8},
		{-9, 10, -11, 12},
	}
	result := solution(exampleMatrix)
	for _, indexPair := range result {
		fmt.Printf("[%d, %d]\n", indexPair[0], indexPair[1])
	}
}
