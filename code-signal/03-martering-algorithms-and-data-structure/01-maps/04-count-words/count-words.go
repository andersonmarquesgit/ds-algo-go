package main

import (
	"fmt"
	"strings"
)

func CountWordsNaive(text string) {
	wordsList := []string{}           // Slice to store unique words
	countList := []int{}              // Corresponding slice to store word counts
	words := strings.Split(text, " ") // Split the text into words using space as delimiter

	for _, word := range words {
		index := indexOf(wordsList, word)
		if index != -1 {
			countList[index]++ // If word exists, increment its count
		} else {
			wordsList = append(wordsList, word) // Add new word to list
			countList = append(countList, 1)    // Initialize its count to 1
		}
	}

	// Print each word and its count
	for i, word := range wordsList {
		fmt.Printf("%s: %d\n", word, countList[i])
	}
}

// Helper function to find index of a word in the slice
func indexOf(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i // Return index if word is found
		}
	}
	return -1 // Return -1 if word is not found
}

func CountWordsOptimized(text string) {
	wordCount := make(map[string]int) // Initialize a map to store word frequencies
	words := strings.Split(text, " ") // Split the text into words using space as delimiter

	for _, word := range words {
		wordCount[word]++ // Increment the counter for each word in the map
	}
	fmt.Println(wordCount) // Print the map, showing each word and its count
}

func main() {
	text := "Go Go Go"
	CountWordsNaive(text)

	CountWordsOptimized(text)
}
