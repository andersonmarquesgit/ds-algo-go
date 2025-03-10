package main

import "fmt"

func FindLastUniqueWordNaive(words []string) string {
	// Traverse the list from the end
	for i := len(words) - 1; i >= 0; i-- {
		isUnique := true
		// Compare the current word to all other words
		for j := 0; j < len(words); j++ {
			// If a duplicate is found, mark as not unique
			if i != j && words[i] == words[j] {
				isUnique = false
				break
			}
		}

		// If the word is unique, return it
		if isUnique {
			return words[i]
		}
	}

	// If no unique word is found, return an empty string
	return ""
}

func FindLastUniqueWordEfficient(words []string) string {
	// Initialize a map to store the count of each word
	wordsMap := make(map[string]int)

	// Traverse the list to populate word counts
	for _, word := range words {
		wordsMap[word]++
	}

	// Iterate from the end to find the last unique word
	for i := len(words) - 1; i >= 0; i-- {
		if wordsMap[words[i]] == 1 {
			return words[i] // Return the last unique word
		}
	}
	return "" // If no unique word is found
}

func main() {
	s1 := []string{"anderson", "maria", "maria"}

	fmt.Println(FindLastUniqueWordNaive(s1))
	fmt.Println(FindLastUniqueWordEfficient(s1))
}
