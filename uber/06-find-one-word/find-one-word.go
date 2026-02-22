package main

import "fmt"

func findWord(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == word[0] {
				if dfs(board, word, row, col, 0, visited) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, row, col, index int, visited [][]bool) bool {
	if index == len(word)-1 {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || visited[row][col] || board[row][col] != word[index] {
		return false
	}

	visited[row][col] = true

	for _, dir := range directions {
		nextRow := row + dir[0]
		nextCol := col + dir[1]

		if dfs(board, word, nextRow, nextCol, index+1, visited) {
			return true
		}
	}

	return false

}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	board := [][]byte{
		{'U', 'B', 'E', 'R'},
		{'S', 'C', 'A', 'S'},
		{'A', 'D', 'T', 'E'},
	}
	word := "UBER"
	fmt.Println(findWord(board, word))
}
