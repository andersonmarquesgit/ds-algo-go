package main

import "fmt"

/*
In this challenge, let's adjust our matrix traversal method! Originally,
our method involved traversing even-indexed columns upwards and odd-indexed columns downwards.

Your task is to reverse this pattern without adding new lines of code.
Make the necessary change in the condition within the loop so that even-indexed columns
are traversed downwards, and odd-indexed ones upwards. Let's enhance this pattern!
*/

func PrintMatrixTraversal(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	for col := cols - 1; col >= 0; col-- {
		if col%2 != 0 {
			for row := rows - 1; row >= 0; row-- {
				fmt.Print(matrix[row][col], " ")
			}
		} else {
			for row := 0; row < rows; row++ {
				fmt.Print(matrix[row][col], " ")
			}
		}
	}
	fmt.Println()
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	PrintMatrixTraversal(matrix)
}
