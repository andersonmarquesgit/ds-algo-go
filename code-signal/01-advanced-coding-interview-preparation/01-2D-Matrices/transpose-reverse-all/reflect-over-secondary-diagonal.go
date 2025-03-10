package main

import (
	"fmt"
)

func ReflectOverSecondaryDiagonal(matrix [][]int) [][]int {
	size := len(matrix)
	cols := len(matrix[0])
	newMatrix := make([][]int, size)
	for i := range newMatrix {
		newMatrix[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		// TODO: Complete the code to obtain the reflected square matrix in newMatrix.
		for j := 0; j < cols; j++ {
			newMatrix[j][cols-i-1] = matrix[i][cols-j-1]
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

	// TODO: Call the function on squareMatrix and store the result in transformedMatrix.
	result := ReflectOverSecondaryDiagonal(squareMatrix)
	// Print the transformed matrix.
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}

}
