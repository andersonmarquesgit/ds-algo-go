package main

import "fmt"

func InterleaveMatrices(matrixA, matrixB [][]int, submatrixCoords [][]int) [][]int {
	// TODO: First we need get the startRowA, endRowA, startRowB, startColB
	startRowA, endRowA := submatrixCoords[0][0]-1, submatrixCoords[0][1]-1
	startColA, endColA := submatrixCoords[0][2]-1, submatrixCoords[0][3]-1
	startRowB := submatrixCoords[1][0] - 1
	startColB := submatrixCoords[1][2] - 1

	rowNum := endRowA - startRowA + 1
	matrixResult := make([][]int, rowNum)

	for i := 0; i < rowNum; i++ { // Count movements we need on rows
		var newRow []int
		for j := 0; j <= endColA-startColA; j++ { // Count movement we need on cols
			newRow = append(newRow, matrixA[startRowA+i][startColA+j])
			newRow = append(newRow, matrixB[startRowB+i][startColB+j])
		}
		matrixResult[i] = newRow
	}

	fmt.Println(matrixResult)

	return [][]int{}
}

func main() {
	matrixA := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	matrixB := [][]int{
		{11, 12, 13},
		{14, 15, 16},
		{17, 18, 19},
	}

	submatrixCoords := [][]int{
		{2, 3, 2, 3},
		{1, 2, 1, 2},
	}

	result := InterleaveMatrices(matrixA, matrixB, submatrixCoords)

	for _, row := range result {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
