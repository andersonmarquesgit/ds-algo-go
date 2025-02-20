package main

import (
	"fmt"
)

func MatrixBoundaryConcatenation(matrixA [][]int, matrixB [][]int, n int) []int {
	// TODO: First we need the directions
	// TODO: We need create a loop for iterate N layers, starting outermost elements and next inner "ring" of elements
	// TODO: We need update limites for each layer, in A and B
	// TODO: Append current element in result and get next row and col for verify limits
	// TODO: If we are in the boundaries is necessary change direction:

	// TODO: We need a result slice
	var result []int

	directions := [][]int{
		{0, 1},  // Right  → (row, col+1)
		{1, 0},  // Down    ↓ (row+1, col)
		{0, -1}, // Left ← (row, col-1)
		{-1, 0}, // Up     ↑ (row-1, col)
	}

	rowsA, colsA := len(matrixA), len(matrixA[0])
	rowsB, colsB := len(matrixB), len(matrixB[0])

	for layer := 0; layer < n; layer++ {
		// Define limits for this layer
		topA, bottomA, leftA, rightA := layer, rowsA-layer-1, layer, colsA-layer-1
		topB, bottomB, leftB, rightB := layer, rowsB-layer-1, layer, colsB-layer-1

		// Percorrer camada da matriz A
		row, col := topA, leftA
		dirIndex := 0 // Starting right direction

		for {
			result = append(result, matrixA[row][col])
			nextRow := row + directions[dirIndex][0]
			nextCol := col + directions[dirIndex][1]

			// Verify limits
			if nextCol > rightA { // Change direction (right → down)
				dirIndex++
				nextRow = row + directions[dirIndex][0]
				nextCol = col + directions[dirIndex][1]
			} else if nextRow > bottomA { // Change direction (down → left)
				dirIndex++
				nextRow = row + directions[dirIndex][0]
				nextCol = col + directions[dirIndex][1]
			} else if nextCol < leftA { // Change direction (left → up)
				dirIndex++
				nextRow = row + directions[dirIndex][0]
				nextCol = col + directions[dirIndex][1]
			} else if nextRow < topA { // Change direction (up → finish layer)
				break
			}

			row, col = nextRow, nextCol
		}

		// Percorrer camada da matriz B
		row, col = topB, leftB
		dirIndex = 0

		for {
			result = append(result, matrixB[row][col])
			nextRow := row + directions[dirIndex][0]
			nextCol := col + directions[dirIndex][1]

			// Verify limits
			if nextCol > rightB { // Change direction (right → down)
				dirIndex++
				nextRow = row + directions[dirIndex][0]
				nextCol = col + directions[dirIndex][1]
			} else if nextRow > bottomB { // Change direction (down → left)
				dirIndex++
				nextRow = row + directions[dirIndex][0]
				nextCol = col + directions[dirIndex][1]
			} else if nextCol < leftB { // Change direction (left → up)
				dirIndex++
				nextRow = row + directions[dirIndex][0]
				nextCol = col + directions[dirIndex][1]
			} else if nextRow < topB { // Change direction (up → finish layer)
				break
			}

			row, col = nextRow, nextCol
		}

	}

	return result
}

func main() {
	matrixA := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrixB := [][]int{
		{17, 18, 19, 20},
		{21, 22, 23, 24},
		{25, 26, 27, 28},
		{29, 30, 31, 32},
	}

	result := MatrixBoundaryConcatenation(matrixA, matrixB, 2)
	fmt.Println(result)
}
