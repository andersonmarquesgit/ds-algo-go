package main

import (
	"fmt"
)

func FindPositions(board [][]rune) [][2]int {
	var positions [][2]int

	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'E' {
				if (i > 0 && board[i-1][j] == 'E') ||
					(i < rows-1 && board[i+1][j] == 'E') ||
					(j > 0 && board[i][j-1] == 'E') ||
					(j < cols-1 && board[i][j+1] == 'E') {
					positions = append(positions, [2]int{i, j})
				}
			}
		}
	}
	return positions
}

func main() {
	board := [][]rune{
		{'P', 'E', 'E', 'P'},
		{'E', 'P', 'E', 'P'},
		{'P', 'E', 'P', 'P'},
		{'P', 'E', 'P', 'E'},
	}

	positions := FindPositions(board)

	for _, pos := range positions {
		fmt.Printf("(%d, %d)\n", pos[0], pos[1])
	}
}
