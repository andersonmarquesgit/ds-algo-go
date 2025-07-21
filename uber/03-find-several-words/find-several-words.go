package main

import "fmt"

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func findSeveralWords(board [][]byte, words []string) map[string]bool {
	rows, cols := len(board), len(board[0])
	wordsVerify := make(map[string]bool)

	// Percorro todas as palavras procuradas
	for _, word := range words {
		// Se a palavra já foi encontrada, não precisa verificar novamente
		if _, exists := wordsVerify[word]; exists {
			continue
		}

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				if board[row][col] == word[0] {
					// Criamos uma matriz `visited` para cada nova busca de palavra
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

	// Marca a célula como visitada
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

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	words := []string{"ABCCED", "SEE", "ABCB"}
	fmt.Println(findSeveralWords(board, words))
}
