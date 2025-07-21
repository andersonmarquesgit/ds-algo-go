package main

import "fmt"

func IsPositionStrategic(board [][]rune, row, col int) bool {
	return (row > 0 && board[row-1][col] == 'E') ||
		(row < len(board)-1 && board[row+1][col] == 'E') ||
		(col > 0 && board[row][col-1] == 'E') ||
		(col < len(board[0])-1 && board[row][col+1] == 'E')
}

func main() {
	// Board game configuration: 'E' for empty, 'P' for piece
	gameBoard := [][]rune{
		{'E', 'P'},
		{'P', 'E'},
	}

	// TODO: Check if position (0, 1) is strategic and print the result
	fmt.Println(IsPositionStrategic(gameBoard, 1, 1))
}
