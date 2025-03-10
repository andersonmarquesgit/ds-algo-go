package main

import (
	"fmt"
	"sort"
)

func RecommendBooks(userHistory []int, popularBooks []int, unpopularBooks []int) []int {
	// TODO: We can store the popular books in map.
	popularSet := make(map[int]bool)
	for _, book := range popularBooks {
		popularSet[book] = true
	}

	// TODO: We iterate over user history and unpopular books
	for _, book := range userHistory {
		delete(popularSet, book) // Se existir, será removido
	}

	for _, book := range unpopularBooks {
		delete(popularSet, book)
	}

	recommend := make([]int, 0, len(popularSet))
	for book := range popularSet {
		recommend = append(recommend, book)
	}
	sort.Ints(recommend) // Ordenação simples

	return recommend
}

func main() {
	userHistory := []int{1, 2, 3, 4, 5}
	popularBooks := []int{1, 2, 3, 6, 7, 8, 9, 10}
	unpopularBooks := []int{4, 5, 11}

	result := RecommendBooks(userHistory, popularBooks, unpopularBooks)

	// Should print: [6, 7, 8, 9, 10]
	fmt.Println(result)
}
