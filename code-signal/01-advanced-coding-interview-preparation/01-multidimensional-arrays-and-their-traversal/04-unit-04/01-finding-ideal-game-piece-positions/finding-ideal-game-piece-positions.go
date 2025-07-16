package main

import "fmt"

/*
function so that it now identifies potential positions for a new game piece based only on the ability to
move left or right on the board. We're narrowing down our criteria to only
consider horizontal movement, still within the paradigm of exploring 2D arrays.
Update the code to reflect this change and observe the altered outcome.

Imagine a 3x3 board like this:

Copy to clipboard
E P E
P E P
E E E
For this board, we want to find positions where a new piece can be placed and can move left or right to another empty cell.

In the first row, (0, 0) and (0, 2) are empty, but they don't have empty neighbors horizontally.
In the second row, (1, 1) is empty, but it doesn't have empty neighbors horizontally.
In the third row, (2, 0), (2, 1), and (2, 2) are all empty, and each has at least one empty neighbor horizontally.
So, the positions where a new piece can be placed are (2, 0), (2, 1), and (2, 2).
*/

func FindPositionsHorizontals(board [][]rune) [][2]int {
	var positions [][2]int
	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'E' {
				if j > 0 && board[i][j-1] == 'E' || j < cols-1 && board[i][j+1] == 'E' {
					positions = append(positions, [2]int{i, j})
				}
			}
		}
	}

	return positions
}

func FindPositions(board [][]rune) [][2]int {
	var positions [][2]int
	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'E' {
				if j > 0 && board[i][j-1] == 'E' ||
					j < cols-1 && board[i][j+1] == 'E' ||
					i > 0 && board[i-1][j] == 'E' ||
					i < rows-1 && board[i+1][j] == 'E' {
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

	fmt.Println(" Horizontal Positions")

	positions := FindPositionsHorizontals(board)

	for _, pos := range positions {
		fmt.Printf("(%d, %d)\n", pos[0], pos[1])
	}

	fmt.Println(" All Positions")

	positions2 := FindPositions(board)

	for _, pos := range positions2 {
		fmt.Printf("(%d, %d)\n", pos[0], pos[1])
	}
}
