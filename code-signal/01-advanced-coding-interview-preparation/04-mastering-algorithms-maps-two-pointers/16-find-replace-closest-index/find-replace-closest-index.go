package main

import (
	"fmt"
	"math"
	"sort"
)

// Busca binária para encontrar o índice do número mais próximo em X
func closestIndex(X []int, target float64) int {
	left, right := 0, len(X)-1
	closestIndex := 0
	minDiff := math.MaxFloat64

	for left <= right {
		mid := left + (right-left)/2
		diff := math.Abs(float64(X[mid]) - target)

		if diff < minDiff {
			minDiff = diff
			closestIndex = mid
		} else if diff == minDiff && mid < closestIndex {
			closestIndex = mid
		}

		// Ajuste da busca binária
		if float64(X[mid]) < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return closestIndex
}

// Função principal para encontrar os valores correspondentes em Y
func FindAndReplace(X []int, Y []int) []int {
	// Criamos uma cópia de X para poder ordená-la sem modificar o original
	sortedX := append([]int{}, X...)
	sort.Ints(sortedX)

	// Mapeamos os valores de X para seus índices originais
	indexMap := make(map[int]int)
	for i, num := range X {
		indexMap[num] = i
	}

	// Criamos um slice para armazenar a resposta final
	result := make([]int, len(Y))

	// Para cada elemento em Y, encontramos o número mais próximo de Y[i] / 2 em X
	for i, num := range Y {
		target := float64(num) / 2
		closestIdxSorted := closestIndex(sortedX, target)         // Índice no X ordenado
		closestOriginalIdx := indexMap[sortedX[closestIdxSorted]] // Índice no X original
		result[i] = Y[closestOriginalIdx]                         // Pegamos o valor de Y correspondente
	}

	return result
}

func main() {
	X := []int{4, 12, 9, 20}
	Y := []int{10, 20, 40, 50}
	result := FindAndReplace(X, Y)
	fmt.Println(result) // Esperado: [10, 40, 50, 50]
}
