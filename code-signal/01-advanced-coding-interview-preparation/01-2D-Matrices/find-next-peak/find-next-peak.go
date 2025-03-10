package main

import "fmt"

// TODO: Implement the FindNextPeak function
func FindNextPeak(matrix [][]int, row, col int) (int, int) {
	// TODO: Define directions
	directions := [][]int{
		{-1, 0},  // UP
		{1, 0},   // DOWN
		{0, -1},  // LEFT
		{0, 1},   // RIGHT
		{-1, -1}, // left UP
		{-1, 1},  // right DOWN
		{1, -1},  // LEFT down
		{1, 1},   // RIGHT down
	}

	rows := len(matrix)
	cols := len(matrix[0])
	rowPeak, colPeak := row, col

	currentHeight := matrix[row][col]

	for _, direction := range directions {
		r := row + direction[0]
		c := col + direction[1]

		// TODO: Fill the if statement
		if r >= 0 && r < rows && c >= 0 && c < cols && currentHeight < matrix[r][c] {
			//return r, c // Return coordinates of the next higher peak
			rowPeak, colPeak = r, c
		}
	}

	//return row, col // No higher peak, stay in place
	return rowPeak, colPeak // Dessa maneira é possível analisar o pico considerando diagonais
}

func main() {
	trailMap := [][]int{
		{3, 5, 6},
		{4, 7, 8},
		{1, 2, 9},
	}

	startPosition := []int{1, 1} // Starting elevation
	nextRow, nextCol := FindNextPeak(trailMap, startPosition[0], startPosition[1])

	fmt.Printf("Next peak at coordinates: [%d, %d]\n", nextRow, nextCol)
}
