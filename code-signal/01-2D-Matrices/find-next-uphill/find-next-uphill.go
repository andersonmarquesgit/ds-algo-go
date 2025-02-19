package main

import "fmt"

func FindNextUphill(grid [][]int, row, col int) *int {
	// Up, down, left, right, left-up, right-up, left-down, right-down - adicionei as direções diagonais
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	nextVal := grid[row][col]
	rows := len(grid)
	cols := len(grid[0])

	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] > nextVal {
			nextVal = grid[newRow][newCol]
		}
	}
	if nextVal != grid[row][col] {
		return &nextVal
	}
	return nil
}

func main() {
	trailGrid := [][]int{
		{1, 2, 3},
		{6, 5, 8},
		{7, 4, 9},
	}
	startRow, startCol := 1, 1
	// Prints the value uphill from the start position or nil if there's no uphill
	nextUphill := FindNextUphill(trailGrid, startRow, startCol)
	if nextUphill != nil {
		fmt.Println(*nextUphill)
	} else {
		fmt.Println("nil")
	}
}
