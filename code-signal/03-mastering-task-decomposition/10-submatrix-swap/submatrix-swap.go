package main

import "fmt"

func SubmatrixSwap(matrix [][]int, coordS1 []int, coordS2 []int) [][]int {
	// TODO: we need the size for submatrix, we can use one of coord (rows and cols)
	// TODO: with size we can create a loop for each and swap the elements
	rows := coordS1[1] - coordS1[0] // Altura da submatriz
	cols := coordS1[3] - coordS1[2] // Largura da submatriz

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentRowS1 := coordS1[0] + i
			currentColS1 := coordS1[2] + j
			currentRowS2 := coordS2[0] + i
			currentColS2 := coordS2[2] + j
			temp := matrix[currentRowS1][currentColS1]
			matrix[currentRowS1][currentColS1] = matrix[currentRowS2][currentColS2]
			matrix[currentRowS2][currentColS2] = temp
		}
	}

	return matrix
}

func main() {
	M := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	coordS1 := []int{0, 2, 2, 4}
	coordS2 := []int{2, 4, 0, 2}

	result := SubmatrixSwap(M, coordS1, coordS2)

	for _, row := range result {
		fmt.Println(row)
	}
}
