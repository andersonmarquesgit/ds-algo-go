package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

func findWord(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == word[0] && dfs(board, word, row, col, 0, visited) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, row, col, wordIndex int, visited [][]bool) bool {
	if wordIndex == len(word) {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || visited[row][col] || board[row][col] != word[wordIndex] {
		return false
	}

	for _, dir := range directions {
		nextRow := row + dir[0]
		nextCol := col + dir[1]
		if dfs(board, word, nextRow, nextCol, wordIndex+1, visited) {
			return true
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(findWord(board, word))
}
