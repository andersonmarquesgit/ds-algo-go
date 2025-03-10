package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to return a sorted string, used to identify anagram signatures
func SortCharacters(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// Function to find anagrams in report2 that are unique and have a matching anagram in report1
func FindAnagrams(report1, report2 []string) int { // O(n * k log k + m * k log k)
	sortedWordsInReport1 := make(map[string]bool)
	// TODO: fill in sortedWordsInReport1
	for _, word1 := range report1 {
		sortedWordsInReport1[SortCharacters(word1)] = true
	}

	anagramsMatched := make(map[string]bool)
	lengthSum := 0

	for _, word := range report2 {
		// TODO: sort the letters in the word
		// TODO: check if word has an anagram match
		// TODO: add to the counter if the match is found for the first time
		if sortedWordsInReport1[SortCharacters(word)] && !anagramsMatched[word] {
			lengthSum += len(word)
			anagramsMatched[word] = true
		}
	}

	return lengthSum
}

func main() {
	report1 := []string{"cat", "dog", "tac", "god", "act"}
	report2 := []string{"tca", "ogd", "atc", "taco"}
	result := FindAnagrams(report1, report2)
	fmt.Println(result) // output: 9

	report3 := []string{"rat", "tar", "bat", "tab", "bats"}
	report4 := []string{"tra", "art", "abr"}
	result2 := FindAnagrams(report3, report4)
	fmt.Println(result2) // output: 6
}
