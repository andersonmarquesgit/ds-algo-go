package main

import (
	"fmt"
)

func EvaluateMove(board [][]rune, row, col int) bool {
	// TODO: Check if a move to the given cell is possible; write a condition to check if the cell is empty.
	// TODO: Check if at least one neighboring cell is empty (not diagonally).
	if board[row][col] == 'E' {
		if (row > 0 && board[row-1][col] == 'E') || //UP
			(row < len(board)-1 && board[row+1][col] == 'E') || //DOWN
			(col > 0 && board[row][col-1] == 'E') || //LEFT
			(row < len(board[0])-1 && board[row][col+1] == 'E') {
			return true
		}
	}

	return false // Placeholder return statement
}

func main() {
	// Assuming a constant 2D slice representing a board
	board := [][]rune{
		{'P', 'E', 'E', 'P'},
		{'E', 'P', 'E', 'P'},
		{'P', 'E', 'P', 'P'},
		{'P', 'E', 'P', 'E'},
	}

	rows := len(board)
	cols := len(board[0])
	// TODO: Write a loop or loops to find all valid move positions and display them.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(EvaluateMove(board, i, j), " ")
		}
		fmt.Println()
	}
}
