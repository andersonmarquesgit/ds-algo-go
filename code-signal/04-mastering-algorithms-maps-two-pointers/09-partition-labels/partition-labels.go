package main

import (
	"fmt"
)

func PartitionLabels(s string) []int {
	// Passo 1: Criar um mapa com a última posição de cada letra na string
	lastOccurrence := make(map[rune]int)
	for i, char := range s {
		lastOccurrence[char] = i
	}

	// Passo 2: Criar as partições
	var partitions []int
	start, end := 0, 0

	for i, char := range s {
		// Expandir a partição até a última ocorrência da letra atual
		if lastOccurrence[char] > end {
			end = lastOccurrence[char]
		}

		// Se chegamos no fim da partição, armazenamos o tamanho e iniciamos outra
		if i == end {
			partitions = append(partitions, end-start+1)
			start = i + 1
		}
	}

	return partitions
}

func main() {
	lengths := PartitionLabels("abacdcd")
	for _, length := range lengths {
		fmt.Print(length, " ")
	}
	// Saída esperada: 3 4
}
