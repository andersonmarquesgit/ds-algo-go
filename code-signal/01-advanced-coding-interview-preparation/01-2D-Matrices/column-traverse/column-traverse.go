package main

import (
	"fmt"
)

func ColumnTraverse(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	row := rows - 1
	col := cols - 1
	direction := "up"
	output := make([]int, rows*cols)
	index := 0

	for index < rows*cols {
		output[index] = matrix[row][col]
		index++
		// TODO: Implement logic to change direction and move left when hitting the top or bottom of the column
		// Otherwise, continue moving in the same direction
		if direction == "up" {
			if row <= 0 {
				direction = "down"
				col--
			} else {
				row--
			}
		} else {
			if row >= rows-1 {
				direction = "up"
				col--
			} else {
				row++
			}
		}
	}

	return output
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	result := ColumnTraverse(matrix)
	for _, num := range result {
		fmt.Printf("%d ", num)
	}
}
