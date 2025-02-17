package main

import "fmt"

// TODO: Fix the function to transpose a seating arrangement represented as a 2D array
func TransposeSeating(seating [][]int) [][]int {
	rows := len(seating)
	cols := len(seating[0])
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = seating[i][j]
		}
	}

	return transposed
}

func main() {
	// Sample restaurant seating arrangement represented as a 2D array
	restaurantSeating := [][]int{
		{10, 11, 12},
		{20, 21, 22},
	}

	// Transpose the seating arrangement
	transposedSeating := TransposeSeating(restaurantSeating)

	for i := 0; i < len(transposedSeating); i++ {
		for j := 0; j < len(transposedSeating[0]); j++ {
			fmt.Print(transposedSeating[i][j], " ")
		}
		fmt.Println()
	}
}
