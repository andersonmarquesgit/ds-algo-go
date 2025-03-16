package main

import "fmt"

// Movimentos possíveis
var directions = [][]int{
	{-1, 0}, // Cima
	{1, 0},  // Baixo
	{0, 1},  // Direita
	{0, -1}, // Esquerda
}

// Função principal que testa se a palavra existe
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Inicia a busca a partir de qualquer célula que corresponda ao primeiro caractere da palavra
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] && dfs(board, word, i, j, 0, visited) {
				return true
			}
		}
	}

	return false
}

/*
📌 Complexidade do Algoritmo
O algoritmo verifica cada palavra separadamente, e para cada palavra, ele pode iniciar uma busca DFS a partir
de qualquer célula da matriz. Assim, a complexidade depende de:

Número de palavras na lista (W)
Número de células na matriz (R * C)
Custo da busca DFS, que pode explorar até 4^L caminhos (L = comprimento da palavra)

🔍 Passo a passo da complexidade
Percorremos a lista de palavras → O(W)
Para cada palavra, exploramos toda a matriz (R × C) → O(W * R * C)
Cada busca DFS pode percorrer até 4^L caminhos (pior caso: explorando todas as direções para cada letra) → O(4^L)

💡 Conclusão da complexidade: O(W * R * C * 4^L)
Onde:

W = número de palavras na lista.
R = número de linhas da matriz.
C = número de colunas da matriz.
L = comprimento da palavra mais longa.
4^L = número máximo de chamadas recursivas no DFS.

📌 Casos Específicos
1️⃣ Melhor Caso (Complexidade Reduzida)
Se uma palavra for encontrada rapidamente, o DFS não precisa explorar todas as células. Isso pode reduzir muito o
número de chamadas.

Complexidade esperada: O(W * R * C) se a maioria das palavras for curta e encontrada rapidamente.
*/
func existList(board [][]byte, words []string) map[string]bool {
	rows, cols := len(board), len(board[0])
	wordsVerify := make(map[string]bool)

	for _, word := range words {
		// Se a palavra já foi encontrada, não precisa verificar novamente
		if _, exists := wordsVerify[word]; exists {
			continue
		}

		found := false
		for i := 0; i < rows && !found; i++ {
			for j := 0; j < cols && !found; j++ {
				if board[i][j] == word[0] {
					// Criamos uma matriz `visited` para cada nova busca de palavra
					visited := make([][]bool, rows)
					for k := range visited {
						visited[k] = make([]bool, cols)
					}
					if dfs(board, word, i, j, 0, visited) {
						wordsVerify[word] = true
						found = true // Se encontramos a palavra, paramos a busca
					}
				}
			}
		}
		if !found {
			wordsVerify[word] = false
		}
	}

	return wordsVerify
}

// Função dsf que explora um caminho até o final antes de retroceder (backtracking). Usa menos memória (pilha recursiva).
func dfs(board [][]byte, word string, row, col, index int, visited [][]bool) bool {
	// Se encontrarmos toda a palavra
	if index == len(word) {
		return true
	}

	// Verificações de limite e se a célula já foi usada
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || visited[row][col] || board[row][col] != word[index] {
		return false
	}

	// Marca a célula como visitada
	visited[row][col] = true

	// Explora todas as direções
	for _, dir := range directions {
		if dfs(board, word, row+dir[0], col+dir[1], index+1, visited) {
			return true
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'C', 'I', 'T', 'Y'},
		{'A', 'V', 'C', 'S'},
		{'T', 'T', 'E', 'E'},
		{'A', 'D', 'E', 'E'},
	}
	word := "CITY"
	fmt.Println(exist(board, word)) // true

	words := []string{"CITY", "CAT"}
	fmt.Println(existList(board, words)) // true, true
}
