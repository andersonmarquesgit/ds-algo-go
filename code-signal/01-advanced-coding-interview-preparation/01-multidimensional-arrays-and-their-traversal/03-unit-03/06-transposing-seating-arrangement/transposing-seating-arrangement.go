package main

import "fmt"

/*
Time to apply what you've learned about transposing matrices!
Imagine you are assisting in redesigning the seating arrangement
of a restaurant to accommodate guests better. Your task is to write Go
code from scratch to take an initial seating arrangement (a 2D splice)
and transpose it. Remember, transposing
switches the rows and columns, creating a new perspective on seating.
*/

/*
Example:
1 2
3 4
5 6

After transposing, it becomes:
1 3 5
2 4 6
*/

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
			// [0][0] = [0][0]
			// [1][0] = [0][1]
			// [0][1] = [1][0]
			// [1][1] = [1][1]

		}
	}

	return transposed
}

func main() {
	// Restaurant seating before transposition (rows)
	seatingBefore := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(TransposeSeating(seatingBefore))
}
