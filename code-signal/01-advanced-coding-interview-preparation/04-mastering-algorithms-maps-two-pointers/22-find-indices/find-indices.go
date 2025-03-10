package main

import (
	"fmt"
	"sort"
)

func findIndices(arrA, arrB []int) []int {
	// TODO: We can use a map for store sum of the each pair in arrA
	// TDOO: We can calculate sum of the each pair in arrB and if we have a key (sum) in map and the index is equal, so we have the result
	// TODO: But if we have two or more result, we need compare the i and j
	sumMap := make(map[int][][2]int)

	// Step 1: Populate map with the sums of all pairs in arrA
	for i := 0; i < len(arrA); i++ {
		for j := i + 1; j < len(arrA); j++ {
			pairwiseSum := arrA[i] + arrA[j]
			sumMap[pairwiseSum] = append(sumMap[pairwiseSum], [2]int{i, j})
		}
	}

	// Step 2: Search for the equal pair sum on arrB
	for i := 0; i < len(arrB); i++ {
		for j := i + 1; j < len(arrB); j++ {
			pairwiseSum := arrB[i] + arrB[j]
			if pairs, exists := sumMap[pairwiseSum]; exists {
				if pairs[0][0] == i && pairs[0][1] == j {
					return []int{i, j}
				}
			}
		}
	}

	return []int{}
}

func findIndices2(arrA, arrB []int) []int {
	sumMap := make(map[int][2]int)

	// 1️⃣ Criamos todas as somas possíveis de pares em arrA e armazenamos no mapa
	for i := 0; i < len(arrA); i++ {
		for j := i + 1; j < len(arrA); j++ {
			pairwiseSum := arrA[i] + arrA[j]
			sumMap[pairwiseSum] = [2]int{i, j} // Apenas um par por soma (primeiro encontrado)
		}
	}

	// 2️⃣ Ordenamos arrB para realizar buscas binárias
	sort.Ints(arrB) // O(N log N)

	// 3️⃣ Percorremos arrB e usamos busca binária para encontrar uma soma correspondente
	for i := 0; i < len(arrB); i++ {
		for j := i + 1; j < len(arrB); j++ {
			pairwiseSum := arrB[i] + arrB[j]
			if indices, exists := sumMap[pairwiseSum]; exists {
				return []int{indices[0], indices[1]}
			}
		}
	}

	return []int{}
}

func findIndicesOptimized(arrA, arrB []int) []int {
	diffs := make(map[int]int)

	for i := 0; i < len(arrA); i++ {
		diff := arrA[i] - arrB[i]
		if j, exists := diffs[-diff]; exists {
			return []int{j, i}
		}
		if _, exists := diffs[diff]; !exists {
			diffs[diff] = i
		}
	}

	return []int{}
}

func main() {
	arrA := []int{2, 5, 1, 4}
	arrB := []int{3, 6, 3, 2}
	indices := findIndices(arrA, arrB)
	fmt.Printf("[%d, %d]\n", indices[0], indices[1]) // Output: [2, 3]

	arrA2 := []int{2, 5, 1, 4}
	arrB2 := []int{3, 6, 3, 2}
	indices2 := findIndices2(arrA2, arrB2)
	fmt.Printf("[%d, %d]\n", indices2[0], indices2[1]) // Output: [2, 3]

	arrA3 := []int{2, 5, 1, 4}
	arrB3 := []int{3, 6, 3, 2}
	indices3 := findIndicesOptimized(arrA3, arrB3)
	fmt.Printf("[%d, %d]\n", indices3[0], indices3[1]) // Output: [2, 3]
}
