package main

import "fmt"

/*
|===========||===========||===========|
| 3 |   | 6 ||   | 5 | 8 || 4 |   |   |
| 5 | 2 |   ||   |   |   ||   |   |   |
|   | 8 | 7 ||   |   |   ||   | 3 | 1 |
|===========||===========||===========|
| 1 |   | 2 || 5 |   |   || 3 | 2 |   | ----> 2 repetindo na linha
| 9 |   |   || 8 | 6 | 3 ||   |   | 5 |
|   | 5 |   ||   | 9 |   || 6 |   |   |
|===========||===========||===========|
|   | 3 |   ||   |   | 8 || 2 | 5 |   |
|   | 1 |   ||   |   |   ||   | 7 | 4 |
|   |   | 5 || 2 |   | 6 ||   |   |   |
|===========||===========||===========|
                       ^ 8 repetindo na coluna
Logo output = FALSE
*/

type Boxes [][]int

func verifySudoku(board [][]int) bool {
	rows := make([]map[int]bool, 9)  // Mapas para verificar números em cada linha
	cols := make([]map[int]bool, 9)  // Mapas para verificar números em cada coluna
	subgrids := [3][3]map[int]bool{} // Matriz de mapas para os subgrids 3x3

	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]bool)
		cols[i] = make(map[int]bool)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			subgrids[i][j] = make(map[int]bool)
		}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := board[row][col]

			if num == 0 {
				continue
			}

			if rows[row][num] {
				fmt.Printf("Número %d já existe na linha %d\n", num, row)
				return false
			}
			rows[row][num] = true

			if cols[col][num] {
				fmt.Printf("Número %d já existe na coluna %d\n", num, col)
				return false
			}
			cols[col][num] = true

			subgridRow := row / 3 // Determina a linha do subgrid (0-2)
			subgridCol := col / 3 // Determina a coluna do subgrid (0-2)
			if subgrids[subgridRow][subgridCol][num] {
				fmt.Printf("Número %d já existe no subgrid (%d, %d)\n", num, subgridRow, subgridCol)
				return false
			}
			subgrids[subgridRow][subgridCol][num] = true
		}
	}

	return true
}

func main() {
	board := [][]int{
		{3, 0, 6, 0, 5, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{1, 0, 2, 5, 0, 0, 3, 2, 0}, // 2 repetido na linha
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{0, 3, 0, 0, 0, 8, 2, 5, 0},
		{0, 1, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 0, 0, 0}, // 8 repetido na coluna
	}

	if verifySudoku(board) {
		fmt.Println("Tabuleiro válido.")
	} else {
		fmt.Println("Tabuleiro inválido.")
	}
}
