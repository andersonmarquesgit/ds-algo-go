package main

import (
	"fmt"
)

func FindPeakAltitude(grid [][]int, startRow, startCol int) int {
	rows := len(grid)
	cols := len(grid[0])
	altitude := grid[startRow][startCol]
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {1, -1}, {-1, -1}, {1, 1}, {-1, 1}}

	for _, direction := range directions {
		newRow := startRow + direction[0]
		newCol := startCol + direction[1]

		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] > altitude {
			altitude = grid[newRow][newCol]
		}
	}
	return altitude
}

func main() {
	mountain := [][]int{
		{1, 2, 3},
		{2, 5, 7},
		{4, 6, 9},
	}

	// number 2 is 7
	fmt.Println(FindPeakAltitude(mountain, 0, 1)) // Should print the altitude of the highest peak reachable

	// number 5 is 9
	fmt.Println(FindPeakAltitude(mountain, 1, 1)) // Should print the altitude of the highest peak reachable

}
