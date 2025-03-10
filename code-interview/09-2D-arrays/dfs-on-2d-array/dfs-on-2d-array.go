package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

func traversalDFS(array2D [][]int) []int {
	values := []int{}

	qtdRows, qtdCols := len(array2D), len(array2D[0])

	// Inicializa a matriz 2D
	seen := make([][]bool, qtdRows) // Slice de slices

	for i := 0; i < qtdRows; i++ {
		seen[i] = make([]bool, qtdCols) // Cada linha é um slice de tamanho de colunas do array original
	}

	dfs(array2D, 0, 0, seen, &values)
	return values
}

func dfs(array2D [][]int, row int, col int, seen [][]bool, values *[]int) {
	if row < 0 || col < 0 || row >= len(array2D) || col >= len(array2D[0]) || seen[row][col] {
		return // Não queremos continuar com essa iteração
	}

	*values = append(*values, array2D[row][col])
	seen[row][col] = true // Marcamos que já analisamos aquela posição na matriz

	for i := 0; i < len(directions); i++ {
		currentDirection := directions[i]
		dfs(array2D, row+currentDirection[0], col+currentDirection[1], seen, values)
	}
}

func main() {
	array2D := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	fmt.Println(traversalDFS(array2D))
}
