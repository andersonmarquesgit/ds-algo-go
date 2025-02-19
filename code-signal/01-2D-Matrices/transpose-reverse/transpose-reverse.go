package main

import (
	"fmt"
)

func TransformMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	// TODO: Modify the loop to transpose the matrix in reverse order
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][cols-j-1]
		}
	}

	return transposed
}

func main() {
	seatingChart := [][]int{
		{101, 102, 103, 104},
		{201, 202, 203, 204},
		{301, 302, 303, 304},
	}

	// TODO: Store the result of TransformMatrix in transposedSeating and print it
	transposedSeating := TransformMatrix(seatingChart)

	for i := 0; i < len(transposedSeating); i++ {
		for j := 0; j < len(transposedSeating[0]); j++ {
			fmt.Print(transposedSeating[i][j], " ")
		}
		fmt.Println()
	}

}
