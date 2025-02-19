package main

import "fmt"

func CountSubmatricesWithE(board [][]rune) int {
	// TODO: Initialize a count variable to keep track of 3x3 submatrices with 'E's in all four corners
	countSubmatrices := 0
	rows := len(board)
	cols := len(board[0])

	// TODO: Use a nested loop to go through each element that can be the top-left corner of a 3x3 submatrix
	// TODO: Check if the current 3x3 submatrix has 'E's in all four corners
	// If it does, increment the count

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'E' &&
				j < cols-2 &&
				i < rows-2 &&
				board[i][j+2] == 'E' && // Top-right
				board[i+2][j] == 'E' && //Bottom-left
				board[i+2][j+2] == 'E' {
				countSubmatrices++
			}
		}
	}

	// TODO: Return the count of submatrices with 'E's in all four corners
	return countSubmatrices
}

func main() {
	// TODO: Define a 2D array 'board' with some 'E's and 'P's, representing empty and piece positions respectively
	board := [][]rune{
		{'E', 'P', 'E', 'P'},
		{'P', 'E', 'P', 'E'},
		{'E', 'P', 'E', 'P'},
		{'P', 'E', 'P', 'E'},
	}

	// TODO: Call the function to count the submatrices and output the result
	fmt.Print(CountSubmatricesWithE(board))
}
