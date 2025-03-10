package main

import "fmt"

// Function to check if a position is strategic
func IsPositionStrategic(board [][]rune, row, col int) bool {
	return ((row > 0 && board[row-1][col] == 'E') || // up
		(row < len(board)-1 && board[row+1][col] == 'E') || // down
		(col > 0 && board[row][col-1] == 'E') || // left
		(col < len(board[0])-1 && board[row][col+1] == 'E')) // right
}

func main() {
	// Board game configuration: 'E' for empty, 'P' for piece
	gameBoard := [][]rune{
		{'E', 'P'},
		{'P', 'E'},
	}

	// TODO: Check if position (0, 1) is strategic and print the result
	fmt.Println(IsPositionStrategic(gameBoard, 0, 1))
}
