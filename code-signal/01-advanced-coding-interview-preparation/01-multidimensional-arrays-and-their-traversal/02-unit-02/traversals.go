package main

import (
	"fmt"
	"log"
)

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

/*
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
*/
func ReverseTraverse2(matrix [][]int) []int { // Menos intuitivo e com controle manual de bordas
	rows := len(matrix)
	cols := len(matrix[0])
	row := rows - 1
	col := cols - 1
	output := make([]int, 0, rows*cols)
	index := 0

	for index < rows*cols {
		output = append(output, matrix[row][col]) // [2][3], [2][2], [2][1], [2][0], [1][3] ...
		index++

		if col-1 < 0 {
			row--
			col = cols - 1
		} else {
			col--
		}
	}

	return output
}

func PathTraverse(grid [][]int, startRow, startCol int) []int {
	rows := len(grid)
	cols := len(grid[0])

	// Check the validity of the input
	if startRow < 0 || startRow >= rows || startCol < 0 || startCol >= cols {
		log.Fatal("Invalid input")
	}

	// Define all four possible directions of movement
	directions := [][]int{
		{1, 0}, {-1, 0}, {0, -1}, {0, 1},
	}

	// Start with the value at the starting cell
	visited := []int{grid[startRow][startCol]}

	for {
		currMax := -1
		nextRow, nextCol := -1, -1

		// Loop over each adjacent cell in all the directions
		for _, dir := range directions {
			newRow := startRow + dir[0]
			newCol := startCol + dir[1]

			// If the new cell is out of the grid boundary, ignore it
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}

			// If the new cell's value is greater than the current maximum
			if grid[newRow][newCol] > currMax {
				// Save it as the next cell to visit
				nextRow = newRow
				nextCol = newCol
				currMax = grid[newRow][newCol]
			}
		}

		// If we don't have any valid cell to visit, break from the loop
		if currMax <= grid[startRow][startCol] {
			break
		}

		// Otherwise, go to the next cell
		startRow = nextRow
		startCol = nextCol

		// Append the cell's value to the result slice
		visited = append(visited, currMax)
	}

	// Return the slice of visited cells' values
	return visited
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(ColumnTraverse(matrix))
	fmt.Println(ReverseTraverse(matrix))
	fmt.Println(ReverseTraverse2(matrix))
	fmt.Println(PathTraverse(matrix, 0, 0))
}
