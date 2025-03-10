package main

import (
	"fmt"
	"sort"
)

func findChocPairs(sweetness []int) [][]int {
	// TODO: Step 01 we can sort the sweetness. Is more efficiently for the two techniques
	sort.Ints(sweetness)
	// TODO: Step 02 we need declare the pointers and result slice of slice
	left := 0
	right := len(sweetness) - 1
	result := [][]int{}

	for left < right {
		total := sweetness[left] + sweetness[right]
		if total == 0 {
			result = append(result, []int{sweetness[right], sweetness[left]})
			left++
			right--
		} else if total < 0 {
			left++
		} else {
			right--
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

// Função otimizada usando HashMap para encontrar pares que somam 0
func findChocPairsOptimized(sweetness []int) [][]int {
	// Mapa para armazenar os números já vistos
	sweetMap := make(map[int]bool)
	result := [][]int{}

	for _, num := range sweetness {
		// Se encontramos o complemento negativo no mapa, adicionamos ao resultado
		if sweetMap[-num] {
			result = append(result, []int{num, -num})
		}
		// Marcamos o número como visto no mapa
		sweetMap[num] = true
	}

	return result
}

func main() {
	sweetness := []int{4, 3, -5, 5, -3, -4}
	result := findChocPairs(sweetness)
	for _, pair := range result {
		fmt.Printf("[%d, %d]\n", pair[0], pair[1])
	}

	fmt.Println()

	sweetness2 := []int{4, 3, -5, 5, -3, -4}
	result2 := findChocPairsOptimized(sweetness2)
	for _, pair := range result2 {
		fmt.Printf("[%d, %d]\n", pair[0], pair[1])
	}
}
