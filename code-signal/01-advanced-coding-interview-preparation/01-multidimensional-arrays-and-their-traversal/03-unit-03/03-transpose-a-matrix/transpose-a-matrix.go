package main

import "fmt"

/*
Your task is to write a Go function named TransformMatrix that will take
a 2D array of integers (representing a matrix) as input and return another 2D array which is the transposed version
of the input matrix. To transpose a matrix means to flip it over its diagonal, turning its rows into columns
and columns into rows.
*/

func TransformMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([][]int, cols)

	// Alocação de memória
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][cols-j-1]
		}
	}

	return transposed
}

func main() {
	originalMatrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// TODO: Call TransformMatrix with originalMatrix and store the result in a variable
	// TODO: Print out the transposed matrix
	fmt.Println(TransformMatrix(originalMatrix))
}
