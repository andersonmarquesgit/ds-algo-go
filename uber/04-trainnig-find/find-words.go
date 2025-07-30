package main

import "fmt"

func findWords(board [][]byte, words []string) map[string]bool {
	rows, cols := len(board), len(board[0])
	wordsVerify := make(map[string]bool)

	for _, word := range words {
		wordsVerify[word] = false
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				if board[row][col] == word[0] {
					visited := make([][]bool, rows)
					for k := range visited {
						visited[k] = make([]bool, cols)
					}

					if dfs(board, word, row, col, 0, visited) {
						wordsVerify[word] = true
					}
				}
			}
		}
	}

	return wordsVerify
}

func dfs(board [][]byte, word string, row, col, wordIndex int, visited [][]bool) bool {
	if wordIndex == len(word) {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || visited[row][col] || board[row][col] != word[wordIndex] {
		return false
	}

	visited[row][col] = true

	for _, dir := range directions {
		nextRow := row + dir[0]
		nextCol := col + dir[1]
		if dfs(board, word, nextRow, nextCol, wordIndex+1, visited) {
			return true
		}
	}

	return false
}

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	board := [][]byte{
		{'U', 'B', 'C', 'E'},
		{'S', 'E', 'C', 'S'},
		{'A', 'R', 'E', 'E'},
	}
	words := []string{"UBER", "UBCE", "TESTE"}
	fmt.Println(findWords(board, words))
}
