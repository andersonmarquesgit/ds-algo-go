package main

import "fmt"

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func traversal(array2D [][]int) []int {
	// Obter o tamanho do array e definir os arrays visited, output e queue
	rows, cols := len(array2D), len(array2D[0])
	var queue [][]int
	output := make([]int, 0)

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	// Definir o primeiro elemento que visitar
	queue = append(queue, []int{0, 0})

	// Iterar sobre a queue
	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:] // Remove o primeiro elemento da queue, ou seja o elemento da vez

		currentRow, currentCol := currentPosition[0], currentPosition[1]
		// Validar se não saímos da borda da matriz ou se o elemento já foi visitado
		if currentRow < 0 || currentRow >= rows || currentCol < 0 || currentCol >= cols || visited[currentRow][currentCol] {
			continue
		}

		// Apos garantir os limites a primeira coisa que queremos é marcar a posição como visitada
		visited[currentRow][currentCol] = true

		// Executamos o trecho principal do código, nesse caso apenas adicionar os valores a um array output
		output = append(output, array2D[currentRow][currentCol])

		// Enquanto houver elementos na queue iteramos em todas as direçoes do elemento corrente para adicionar a queue
		for _, dir := range directions {
			nextRol, nextCol := currentRow+dir[0], currentCol+dir[1]
			queue = append(queue, []int{nextRol, nextCol})
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

	fmt.Println(traversal(array2D))
}
