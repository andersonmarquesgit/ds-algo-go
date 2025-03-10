package main

import "fmt"

var directions = [][]int{
	{-1, 0}, //UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

func numberOfIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	islandCount := 0

	// Percorremos toda a matriz
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 { // Encontramos uma ilha
				islandCount++
				dfs(grid, row, col) // Afunda a ilha alterando `1` para `0`
			}
		}
	}

	return islandCount
}

// DFS para "afundar" a ilha e evitar contagem repetida
func dfs(grid [][]int, row, col int) {
	// Verificamos se estamos fora dos limites ou se a célula já é água (`0`)
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == 0 {
		return
	}

	// "Afundamos" a ilha marcando `1` como `0`
	grid[row][col] = 0

	// Chamamos recursivamente para os vizinhos
	for _, dir := range directions {
		dfs(grid, row+dir[0], col+dir[1])
	}
}

func main() {
	array2D := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},
	}

	fmt.Println(numberOfIslands(array2D))

	grid2 := [][]int{
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
	}

	fmt.Println(numberOfIslands(grid2))

	grid3 := [][]int{
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 0},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
	}

	fmt.Println(numberOfIslands(grid3))

	grid4 := [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 1},
	}

	fmt.Println(numberOfIslands(grid4))
}
