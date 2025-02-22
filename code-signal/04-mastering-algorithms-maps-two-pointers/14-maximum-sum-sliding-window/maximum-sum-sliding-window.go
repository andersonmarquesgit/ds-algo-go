package main

import (
	"fmt"
)

func maximumSum(numbers []int, k int) []int {
	if len(numbers) < k {
		return []int{0, -1} // Retorna um índice inválido caso o array seja menor que k
	}

	// Passo 1: Inicializar a soma da primeira janela de tamanho k
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += numbers[i]
	}

	maxSum := currentSum
	startIndex := 0

	// Passo 2: Deslizar a janela pelo array
	for i := k; i < len(numbers); i++ {
		// Remover o primeiro elemento da janela e adicionar o próximo
		currentSum = currentSum - numbers[i-k] + numbers[i]

		// Atualizar o máximo encontrado
		if currentSum > maxSum {
			maxSum = currentSum
			startIndex = i - k + 1
		}
	}

	return []int{maxSum, startIndex}
}

func main() {
	result1 := maximumSum([]int{2, 1, 5, 1, 3, 2}, 3)
	fmt.Printf("[%d, %d]\n", result1[0], result1[1]) // Expected output: [9, 2]

	result2 := maximumSum([]int{2, 1, 5, 1, 3, 2}, 2)
	fmt.Printf("[%d, %d]\n", result2[0], result2[1]) // Expected output: [6, 1]
}
