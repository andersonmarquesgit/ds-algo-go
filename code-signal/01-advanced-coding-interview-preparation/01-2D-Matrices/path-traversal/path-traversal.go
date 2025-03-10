package main

import (
	"fmt"
	"log"
)

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
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := PathTraverse(grid, 1, 1)
	for _, val := range res {
		fmt.Print(val, " ")
	}
	fmt.Println()

	// Outputs: 5 8 9
}
