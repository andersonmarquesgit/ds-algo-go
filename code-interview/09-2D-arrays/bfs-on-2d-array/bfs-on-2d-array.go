package main

import "fmt"

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// T: O(n), S O(n)
func traversalBFS(array2D [][]int) []int {
	rows, cols := len(array2D), len(array2D[0])
	var queue [][]int
	output := make([]int, 0)

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	queue = append(queue, []int{0, 0})
	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:]

		row, col := currentPosition[0], currentPosition[1]

		if row < 0 || col < 0 || row >= rows || col >= cols || visited[row][col] {
			continue
		}

		// Apos garantir os limites a primeira coisa que queremos é marcar a posição como visitada
		visited[row][col] = true

		// Logo depois fazer nossa lógica principal. Nesse caso adicionar no output
		output = append(output, array2D[row][col])

		// Agora podemos iterar em todas as direções do elemento
		for _, direction := range directions {
			nextRow, nextCol := row+direction[0], col+direction[1]
			queue = append(queue, []int{nextRow, nextCol})
		}
	}

	return output
}

func main() {
	array2D := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	fmt.Println(traversalBFS(array2D))
}
