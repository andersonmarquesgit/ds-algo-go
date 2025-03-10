package main

import (
	"fmt"
)

func FindInfluencer(connections []struct{ A, B string }) string {
	// First: Clarifying:
	// We can receive the same connection? For example: {"Alice", "Bob"} and {"Bob", "Alice"}

	// Always receive a connection or we can receive connections nil?
	// TODO: My first approach is a iterate over connections get the direct connections from A to B
	firstConnections := make(map[string][]string)

	for _, conection := range connections {
		firstConnections[conection.A] = append(firstConnections[conection.A], conection.B)
		firstConnections[conection.B] = append(firstConnections[conection.B], conection.A)
	}

	influenceCount := make(map[string]int)

	// Iteramos sobre cada pessoa e somamos as conexões indiretas
	for person, directConnections := range firstConnections {
		visited := make(map[string]bool)
		visited[person] = true // Marca a própria pessoa como visitada

		// Marca conexões diretas como visitadas para excluí-las da contagem
		for _, direct := range directConnections {
			for _, secondDegree := range firstConnections[direct] {
				if !visited[secondDegree] { // Apenas conexões indiretas contam
					influenceCount[person]++
					visited[secondDegree] = true
				}
			}
		}
	}

	// Encontramos a pessoa com maior influência
	maxInfluence := 0
	influencer := ""

	for person, count := range influenceCount {
		// Se for maior, atualizamos
		if count > maxInfluence {
			maxInfluence = count
			influencer = person
		} else if count == maxInfluence && person < influencer {
			// Desempate por ordem alfabética
			influencer = person
		}
	}

	return influencer
}

func main() {
	connections := []struct{ A, B string }{
		{"Alice", "Bob"}, {"Bob", "Charlie"}, {"Alice", "Charlie"},
		{"Alice", "David"}, {"David", "Eve"}, {"Eve", "Frank"},
		{"Bob", "Frank"},
	}
	fmt.Println(FindInfluencer(connections)) // Expected output: 'Alice'
}
