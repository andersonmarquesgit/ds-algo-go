package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

func numberOfIslands(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	numberOfIslands := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				numberOfIslands++
				dfs(grid, i, j)
				//bfs(grid, i, j)
			}
		}
	}

	return numberOfIslands
}

func dfs(grid [][]int, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == 0 {
		return
	}

	grid[row][col] = 0

	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := col + direction[1]

		dfs(grid, nextRow, nextCol)
	}
}

func bfs(grid [][]int, row, col int) {
	queue := [][2]int{{row, col}}
	grid[row][col] = 0

	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:]

		currentRow, currentCol := currentPosition[0], currentPosition[1]
		for _, direction := range directions {
			nextRow := currentRow + direction[0]
			nextCol := currentCol + direction[1]

			if nextRow >= 0 && nextCol >= 0 && nextRow < len(grid) && nextCol < len(grid[0]) && grid[nextRow][nextCol] == 1 {
				grid[nextRow][nextCol] = 0
				queue = append(queue, [2]int{nextRow, nextCol})
			}
		}
	}
}

func main() {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	fmt.Println(numberOfIslands(grid))

	grid2 := [][]int{
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
	}

	fmt.Println(numberOfIslands(grid2))

}

/*
### Complexidade
- **Tempo**: O(m × n), onde m é o número de linhas e n é o número de colunas
- **Espaço**: O(m × n) no pior caso devido à recursão do DFS

*/

/*
### Conclusão
Para este problema específico de contar ilhas, a escolha entre DFS e BFS é principalmente uma questão de preferência, pois:
1. Ambos visitam cada célula exatamente uma vez
2. Ambos têm a mesma complexidade de tempo O(m × n)
3. O resultado final será o mesmo

Eu manteria a implementação DFS atual porque:
1. É mais simples de entender e manter
2. O problema não requer encontrar caminhos mais curtos
3. As matrizes típicas não são grandes o suficiente para causar estouro de pilha
4. A memória extra usada pela recursão não é uma preocupação significativa neste caso

*/
