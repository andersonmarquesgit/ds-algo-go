package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// BFS a partir da borda
	for r := 0; r < rows; r++ {
		if board[r][0] == 'O' {
			bfs(board, r, 0)
		}
		if board[r][cols-1] == 'O' {
			bfs(board, r, cols-1)
		}
	}

	for c := 0; c < cols; c++ {
		if board[0][c] == 'O' {
			bfs(board, 0, c)
		}
		if board[rows-1][c] == 'O' {
			bfs(board, rows-1, c)
		}
	}

	// Pós-processamento
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X' // cercado
			} else if board[r][c] == 'E' {
				board[r][c] = 'O' // seguro
			}
		}
	}
}

func bfs(board [][]byte, row, col int) {
	queue := [][]int{{row, col}}
	board[row][col] = 'E' // marca como seguro (escaped)

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			r := pos[0] + d[0]
			c := pos[1] + d[1]

			if r >= 0 && c >= 0 && r < len(board) && c < len(board[0]) && board[r][c] == 'O' {
				board[r][c] = 'E'
				queue = append(queue, []int{r, c})
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	fmt.Println(board)
	solve(board)
	fmt.Println(board)
}
