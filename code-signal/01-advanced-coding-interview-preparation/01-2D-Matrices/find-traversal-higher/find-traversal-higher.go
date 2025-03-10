package main

import (
	"fmt"
	"log"
)

func TrekPath(elevationMap [][]int, startRow int, startCol int) []int {
	rows := len(elevationMap)
	cols := len(elevationMap[0])

	// Check the validity of the input
	if startRow < 0 || startRow >= rows || startCol < 0 || startCol >= cols {
		log.Fatal("Invalid input")
	}

	// TODO: Initialize the path with the starting position's elevation.
	directions := [][]int{
		{-1, 0}, // UP
		{1, 0},  // DOWN
		{0, -1}, // LEFT
		{0, 1},  // RIGHT
	}

	visited := []int{elevationMap[startRow][startCol]}
	currentHigher := elevationMap[startRow][startCol]

	// TODO: Implement the loop to find the path through higher elevations.
	for {
		nextRow, nextCol := -1, -1

		for _, dir := range directions {
			newRow := startRow + dir[0]
			newCol := startCol + dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && currentHigher < elevationMap[newRow][newCol] {
				currentHigher = elevationMap[newRow][newCol]
				nextRow = newRow
				nextCol = newCol
			}
		}

		// If we don't have any valid cell to visit, break from the loop
		if currentHigher <= elevationMap[startRow][startCol] {
			break
		}

		startRow = nextRow
		startCol = nextCol

		visited = append(visited, currentHigher)
	}

	return visited
}

func main() {
	mountain := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 5, 6},
	}
	result := TrekPath(mountain, 1, 1)
	for _, height := range result {
		fmt.Printf("%d ", height)
	}
}
