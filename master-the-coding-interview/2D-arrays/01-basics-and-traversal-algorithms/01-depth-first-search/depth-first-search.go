package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

func traversal(array2D [][]int) []int {
	rows := len(array2D)
	cols := len(array2D[0])
	visited := make([][]bool, rows)
	output := []int{}

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	dfs(array2D, visited, 0, 0, &output)
	return output
}

func dfs(array2D [][]int, visited [][]bool, row, col int, output *[]int) {
	if row < 0 || row >= len(array2D) || col < 0 || col >= len(array2D[0]) || visited[row][col] {
		return
	}

	visited[row][col] = true
	*output = append(*output, array2D[row][col])

	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := col + direction[1]
		dfs(array2D, visited, nextRow, nextCol, output)
	}
}

func main() {
	array2D := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(traversal(array2D))
}
