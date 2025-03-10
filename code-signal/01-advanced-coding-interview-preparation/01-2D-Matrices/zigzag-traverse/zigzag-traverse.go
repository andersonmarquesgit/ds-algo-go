package main

import "fmt"

func ZigzagTraverse(grid [][]int) []int {
	// TODO: Determine the number of rows and columns in 'grid'
	rows := len(grid)
	cols := len(grid[0])

	row := rows - 1
	col := cols - 1

	gridSize := rows * cols

	traversalPath := []int{}

	// TODO: Traverse 'grid' in a zigzag pattern starting from the bottom right
	direction := "up"

	for i := 0; i < gridSize; i++ {
		traversalPath = append(traversalPath, grid[row][col])

		if direction == "up" {
			if row <= 0 {
				direction = "down"
				col--
			} else {
				row--
			}
		} else {
			if row >= rows-1 {
				direction = "up"
				col--
			} else {
				row++
			}

		}
	}

	return traversalPath
}

func main() {
	// TODO: Define a 2x4 'grid' with unique numbers representing items
	grid := [][]int{
		{101, 102, 103, 104},
		{201, 202, 203, 204},
	}

	// TODO: Print the zigzag traversal path of items in 'grid'
	fmt.Print(ZigzagTraverse(grid))
}
