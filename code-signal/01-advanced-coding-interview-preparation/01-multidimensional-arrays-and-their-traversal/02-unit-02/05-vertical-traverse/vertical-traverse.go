package main

import "fmt"

func VerticalTraverse(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	row := rows - 1
	col := cols - 1
	output := make([]int, 0, rows*cols)
	index := 0

	for index < rows*cols {
		output = append(output, matrix[row][col])
		index++

		if row-1 < 0 {
			col--
			row = rows - 1
		} else {
			row--
		}
	}

	return output
}

func main() {
	data := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	traversal := VerticalTraverse(data)
	for _, num := range traversal {
		fmt.Print(num, " ")
	}
}
