/*
Task Statement
Suppose we have a matrix where each cell represents a distinct symbol or integer. Our task is to decode this matrix by
reading the cells in a particular order.

The decoding begins from the top-left cell of the matrix. We move in a bottom-left downward diagonal direction until
we hit the left boundary. Upon hitting the left boundary, we move one cell down (unless we're at the bottom-left corner
already, in which case we move one cell to the right) and start moving in an upward diagonal direction toward
the upper-right.

While moving diagonally up-right, if we hit the top boundary, we move one cell to the right and start moving in a
downward diagonal direction toward the bottom-left. However, if we hit the right boundary while moving diagonally upward,
we move one cell down and start moving in a bottom-left direction. In other words, we keep zigzagging diagonally
across the matrix until every cell in the matrix is visited.

Upon completing this zigzag traversal, we will have a list of traversed cell values. Next, we process this list to
uncover the indices of the perfect square numbers. The function should implement this traversal and return a list
containing the 1-indexed positions of perfect square numbers in the traversed sequence.

{
    {1, 2, 3, 4},
    {5, 6, 7, 8},
    {9, 10, 11, 12}
}

Upon completing the diagonal traversal, we'll get the list: {1, 5, 2, 3, 6, 9, 10, 7, 4, 8, 11, 12}.
From this list, we see that 1, 9, and 4 are perfect squares and are located at the 1st, 6th,
and 9th positions (1-indexed) in the list. Thus, our function returns: {1, 6, 9}.
*/

/*
Step Overview
To tackle this problem, we will take the following steps:

Traverse the Matrix Diagonally: Begin from the top-left cell and zigzag through the matrix in a specific diagonal pattern.
Record Traversed Values: As you traverse, collect the cell values in a list.
Identify Perfect Squares: Inspect the collected values to determine which are perfect squares.
Return Positions: Gather the 1-indexed positions of the perfect squares in the list and return them.
*/

/*
Traversing the Matrix Diagonally
To traverse the matrix diagonally in a zigzag pattern, we need to adhere to a specific movement strategy that alternates
between two directions: down-left and up-right. Here's a clearer breakdown of the traversal logic.

Initialize Variables: We start by initializing row and col to zero, representing the top-left position of the matrix.
We also set dir to 1 for the initial down-left direction.

Traverse the Matrix: We loop through each cell in the matrix using a loop that runs rows * cols times, which is the
total number of elements in the matrix.

Add Current Value: In each iteration, add the current matrix element located at (row, col) to the traversal slice.

Determine Move Direction:

Direction = 1 (down-left):
If row == rows - 1 (bottom boundary), move one step right (col += 1) and change direction to up-right (dir = -1).
If col == 0 (left boundary), move one step down (row += 1) and change direction to up-right (dir = -1).
Otherwise, move diagonally down-left (row += 1; col -= 1).
Direction = -1 (up-right):
If col == cols - 1 (right boundary), move one step down (row += 1) and change direction to down-left (dir = 1).
If row == 0 (top boundary), move one step right (col += 1) and change direction to down-left (dir = 1).
Otherwise, move diagonally up-right (row -= 1; col += 1).
This approach ensures that the traversal covers all cells in a zigzag manner while respecting the matrix boundaries.
*/

package main

import (
	"fmt"
	"math"
)

func DiagonalTraverseAndSquares(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	traversal := []int{}
	results := []int{}

	row, col, dir := 0, 0, 1
	for i := 0; i < rows*cols; i++ {
		traversal = append(traversal, matrix[row][col]) // Append current cell value to traversal slice

		// Determine next position based on current direction
		if dir == 1 { // Moving down-left
			if row == rows-1 { // Hit bottom boundary
				col += 1
				dir = -1 // Switch direction to up-right
			} else if col == 0 { // Hit left boundary
				row += 1
				dir = -1 // Switch direction to up-right
			} else {
				row += 1
				col -= 1 // Continue down-left
			}
		} else { // Moving up-right
			if col == cols-1 { // Hit right boundary
				row += 1
				dir = 1 // Switch direction to down-left
			} else if row == 0 { // Hit top boundary
				col += 1
				dir = 1 // Switch direction to down-left
			} else {
				row -= 1
				col += 1 // Continue up-right
			}
		}
	}

	for idx, val := range traversal {
		root := math.Sqrt(float64(val))
		if root == float64(int(root)) { // Check if the value is a perfect square number.
			results = append(results, idx+1)
		}
	}

	return results
}

func main() {
	matrix := [][]int{
		{16, 2, 3, 13},
		{5, 11, 10, 8},
		{9, 7, 6, 12},
		{4, 14, 15, 1},
	}

	result := DiagonalTraverseAndSquares(matrix)
	fmt.Println(result) // Output: [1, 6, 7, 16]

	/*
		We define a 4x4 matrix matrix.
		We call the DiagonalTraverseAndSquares function with matrix as its argument.
		The function returns a slice [1, 6, 7, 16], which represents the 1-indexed positions
		of perfect square numbers in the traversed sequence.
		For further verification, let's decode this traversal step by step:

		Diagonal traversal sequence: {16, 5, 2, 3, 11, 9, 4, 7, 10, 13, 8, 6, 14, 15, 12, 1}
		Perfect squares in the sequence: 16 (1st), 9 (6th), 4 (7th), 1 (16th)
		Thus, the 1-indexed positions of perfect squares are correctly identified as [1, 6, 7, 16].
	*/
}
