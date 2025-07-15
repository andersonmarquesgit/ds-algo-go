package main

import "fmt"

func ColumnTraverse(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	row := rows - 1
	col := cols - 1
	output := make([]int, 0, rows*cols)
	index := 0
	dir := "up"

	for index < rows*cols {
		output = append(output, matrix[row][col])
		index++

		if dir == "up" {
			if row-1 < 0 {
				dir = "down"
				col--
			} else {
				row--
			}
		} else {
			if row+1 >= rows {
				dir = "up"
				col--
			} else {
				row++
			}
		}
	}

	return output
}

func ReverseTraverse(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	output := make([]int, 0, rows*cols)
	index := 0

	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			output = append(output, matrix[row][col])
			index++
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

	fmt.Println(ColumnTraverse(matrix))
	fmt.Println(ReverseTraverse(matrix))
}
