package main

import "fmt"

func ReflectOverSecondaryDiagonal(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	newMatrix := make([][]int, rows)

	for i := range newMatrix {
		newMatrix[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newMatrix[i][j] = matrix[rows-j-1][cols-i-1]
			// [0][0] = [2][2]
			// [0][1] = [1][2]
			// [0][2] = [0][2]
			// ...
			// [2][2] = [0][0]
		}
	}

	return newMatrix
}

func main() {
	// Example square matrix to reflect over the secondary diagonal
	squareMatrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(ReflectOverSecondaryDiagonal(squareMatrix))

}
