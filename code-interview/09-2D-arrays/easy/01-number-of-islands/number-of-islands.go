package main

import "fmt"

var directions = [][]int{
	{-1, 0}, //UP
	{0, 1},  // RIGHT
	{-1, 0}, // DOWN
	{0, -1}, // LEFT
}

func numberOfIslands(array2D [][]int) int {
	if len(array2D) == 0 {
		return 0
	}

	islandCount := 0

	// We have 3 problemas:
	// 0 -> Do not
	// 1 -> Is a island or part of the island (old)
	/*
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},

		sempre que encontrarmos uma terra podemos converter para água a fim de não usá-lo novamente
	*/

	for row := 0; row < len(array2D); row++ {
		for col := 0; col < len(array2D[0]); col++ {
			if array2D[row][col] == 1 {
				islandCount++
				array2D[row][col] = 0 // sempre que encontrarmos uma terra podemos converter para água a fim de não usá-lo novamente
				queue := [][]int{}
				queue = append(queue, []int{row, col})

				for len(queue) > 0 {
					currentPos := queue
					currentRow := currentPos[0][0]
					currentCol := currentPos[0][1]

					queue = queue[1:]

					for i := 0; i < len(directions); i++ {
						currentDir := directions[i]
						nextRow := currentRow + currentDir[0]
						nextCol := currentCol + currentDir[1]

						if nextRow < 0 || nextRow >= len(array2D) || nextCol < 0 || nextCol >= len(array2D[0]) {
							continue
						}

						if array2D[nextRow][nextCol] == 1 {
							queue = append(queue, []int{row, col})
							array2D[nextRow][nextCol] = 0
						}

					}
				}
			}
		}
	}

	return islandCount
}

func main() {
	array2D := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},
	}

	fmt.Println(numberOfIslands(array2D))
}
