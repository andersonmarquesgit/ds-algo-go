package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortCharacters(word string) string {
	chars := []rune(strings.ToLower(word))
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func FindAnagramsNaive(array1, array2 []string) []string {
	result := []string{}
	// Compare every word in array1 with every word in array2
	for _, word1 := range array1 {
		for _, word2 := range array2 {
			// If sorted characters match, they are anagrams
			if SortCharacters(word1) == SortCharacters(word2) {
				// Avoid duplicates in the result
				if !contains(result, word1) {
					result = append(result, word1)
				}
				break
			}
		}
	}
	return result
}

func contains(s []string, e string) bool {
	// Check for the presence of a word in the result slice
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func FindAnagramsEfficient(array1, array2 []string) []string {
	// Create a map to store sorted signatures of words in array2
	sortedWordsInArray2 := make(map[string]bool)
	for _, word := range array2 {
		sortedWordsInArray2[SortCharacters(word)] = true
	}

	result := []string{}
	// Track matched anagrams to avoid duplicates
	anagramsMatched := make(map[string]bool)
	for _, word := range array1 {
		sortedWord := SortCharacters(word)
		// Check for anagram match and ensure no duplicate in result
		if sortedWordsInArray2[sortedWord] && !anagramsMatched[word] {
			result = append(result, word)
			anagramsMatched[word] = true
		}
	}

	return result
}

func main() {
	array1, array2 := []string{"amor", "listen", "cadeira"}, []string{"roma", "silent", "mesa"}

	fmt.Println(FindAnagramsNaive(array1, array2))
	fmt.Println(FindAnagramsEfficient(array1, array2))
}
