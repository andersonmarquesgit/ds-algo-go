package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	numberOfIslands := 0
	rows, cols := len(grid), len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				numberOfIslands++
				//dfs(grid, row, col)
				bfs(grid, row, col)
			}
		}
	}

	return numberOfIslands
}

func dfs(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'

	for _, dir := range directions {
		nextRow := row + dir[0]
		nextCol := col + dir[1]

		dfs(grid, nextRow, nextCol)
	}
}

func bfs(grid [][]byte, row, col int) {
	queue := [][]int{{row, col}}
	grid[row][col] = '0'

	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:]

		currentRow, currentCol := currentPosition[0], currentPosition[1]
		for _, direction := range directions {
			nextRow := currentRow + direction[0]
			nextCol := currentCol + direction[1]

			if nextRow >= 0 && nextCol >= 0 && nextRow < len(grid) && nextCol < len(grid[0]) && grid[nextRow][nextCol] == '1' {
				grid[nextRow][nextCol] = '0'
				queue = append(queue, []int{nextRow, nextCol})
			}
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid))

	grid2 := [][]byte{
		{'1', '0', '1', '0', '1'},
		{'0', '1', '0', '1', '0'},
		{'1', '0', '1', '0', '1'},
		{'0', '1', '0', '1', '0'},
	}

	fmt.Println(numIslands(grid2))

}
