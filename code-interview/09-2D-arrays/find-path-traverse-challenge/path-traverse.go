package main

import (
	"fmt"
)

// TODO: Create a function PathTraverse
func PathTraverse(grid [][]int, startRow, startCol int) []int {
	rows := len(grid)
	cols := len(grid[0])

	// verify the start position
	if startRow < 0 || startRow >= rows || startCol < 0 || startCol >= cols {
		return nil
	}

	// TODO: Define a 2D slice of slices of integers named 'directions' to represent movement options: up, down, left, and right.
	directions := [][]int{
		{-1, 0}, // UP
		{1, 0},  // DOWN
		{0, -1}, // LEFT
		{0, 1},  // RIGHT
	}

	// TODO: Initialize a slice named 'visited' to store the path values starting with the initial position value.
	currentHeigher := grid[startRow][startCol]
	visited := []int{currentHeigher}

	// TODO: Implement a loop that continues running while movement to a higher adjacent value is possible.
	for {
		nextRow, nextCol := -1, -1

		// TODO: Inside the loop, iterate over each direction to check for valid moves: within bounds and to a higher cell value than the current position.
		for _, dir := range directions {
			newRow := startRow + dir[0]
			newCol := startCol + dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && currentHeigher < grid[newRow][newCol] {
				currentHeigher = grid[newRow][newCol]
				nextRow = newRow
				nextCol = newCol
			}
		}

		// If we don't have any valid cell to visit, break from the loop
		if currentHeigher <= grid[startRow][startCol] {
			break
		}

		startRow = nextRow
		startCol = nextCol

		visited = append(visited, currentHeigher)
	}

	return visited
}

func main() {
	// TODO: Create a 2D array named 'grid' representing ascending values, akin to increasing elevation while hiking up a mountain.
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// TODO: Set the starting position on the grid using startRow and startCol.
	startRow, startCol := 1, 1

	// TODO: Call the PathTraverse function to find the path from the starting point.
	result := PathTraverse(grid, startRow, startCol)

	// TODO: Output the values of the path taken from the starting point.
	for _, height := range result {
		fmt.Printf("%d ", height)
	}
}
