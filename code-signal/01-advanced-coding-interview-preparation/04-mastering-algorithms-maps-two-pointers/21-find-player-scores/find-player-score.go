package main

import (
	"fmt"
)

func GetScores(playerIds, playerScores, queries []int) []int {
	// TODO: We need a map for the store the playerID index position
	// Example: playerIds {1, 2, 3, 4, 5} => map{1:0; 2:1;... 5:0} => For this we can store iterating over queries
	// Criamos um mapa para armazenar os índices dos jogadores
	playerIdx := make(map[int]int)

	// Preenchemos o mapa associando playerId -> índice no array playerScores
	for idx, player := range playerIds {
		playerIdx[player] = idx
	}

	// Criamos um slice para armazenar os resultados na ordem correta das queries
	result := make([]int, 0, len(queries))

	// Iteramos sobre queries para buscar os scores corretos
	for _, query := range queries {
		if idx, exists := playerIdx[query]; exists {
			result = append(result, playerScores[idx])
		}
	}

	return result
}

func main() {
	playerIds := []int{1, 2, 3, 4, 5}
	playerScores := []int{100, 200, 150, 50, 300}
	queries := []int{2, 5, 1}
	result := GetScores(playerIds, playerScores, queries)
	fmt.Println(result)
}
