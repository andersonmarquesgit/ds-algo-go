package main

import (
	"fmt"
	"math"
)

func CalculateMinDistances(wordList []string) map[string]int {
	// Clarify requirements and test cases:
	// {"dog", "cat", "bird", "cat", "dog", "elephant", "dog"} => {dog: 2, cat: 2}
	// {"dog", "cat", "bird", "cat", "dog", "elephant", "elephant", "dog"} => {elephant: 1}
	// {} => nil
	// {"elephant"} => nil

	if len(wordList) <= 1 {
		return nil
	}

	// Mapas para armazenar as últimas ocorrências e as menores distâncias
	lastPosition := make(map[string]int)
	minDistance := make(map[string]int)

	// TODO: We count the animal in wordList and update distance in map
	for idx, word := range wordList {
		if lastIdx, exists := lastPosition[word]; exists {
			// Calcula a distância entre ocorrências consecutivas
			distance := idx - lastIdx
			if currentMin, found := minDistance[word]; found {
				// Atualiza a distância mínima se encontrarmos uma menor
				if distance < currentMin {
					minDistance[word] = distance
				}
			} else {
				// Primeira vez que encontramos um par da palavra
				minDistance[word] = distance
			}
		}
		// Atualiza a última posição da palavra
		lastPosition[word] = idx
	}

	// Se uma palavra aparece apenas uma vez, ela deve ser removida do resultado
	for word, dist := range minDistance {
		if dist == math.MaxInt {
			delete(minDistance, word)
		}
	}

	return minDistance
}

func main() {
	exampleInput := []string{"dog", "cat", "bird", "cat", "dog", "elephant", "dog"}
	result := CalculateMinDistances(exampleInput)
	// Expected output:
	// dog: 2
	// cat: 2
	// Note: The exact order of output pairs might differ.
	for k, v := range result {
		fmt.Printf("%s: %d\n", k, v)
	}
}
