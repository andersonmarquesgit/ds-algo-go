package main

import "fmt"

func IsEmptyNeighbor(board [][]rune, x int, y int) bool {
	rows := len(board)
	cols := len(board[0])

	// Check that (x, y) is within board boundaries
	if x >= 0 && x < rows && y >= 0 && y < cols {
		return board[x][y] == 'E'
	}
	return false
}

func FindPositions(board [][]rune) [][2]int {
	var positions [][2]int
	rows := len(board)
	cols := len(board[0])

	// TODO: Modify the code to check for empty spot with an empty neighbor to the left and right only
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'E' {
				// TODO: Update the condition to check only horizontal neighbors
				if j > 0 && board[i][j-1] == 'E' || j < cols-1 && board[i][j+1] == 'E' {
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
