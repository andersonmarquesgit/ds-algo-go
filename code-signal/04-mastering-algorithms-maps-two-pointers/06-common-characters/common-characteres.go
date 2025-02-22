package main

import (
	"fmt"
	"sort"
)

func CommonCharacters(s string, letters []rune) []rune {
	// TODO: First we need convert the s in map
	set1 := make(map[rune]bool)
	commom := []rune{}

	for _, char := range s {
		set1[char] = true
	}

	// TODO: Second we can iterate over lettters
	for _, letter := range letters {
		// TODO: And we can verify commo character and store in commo slice []
		if _, exists := set1[letter]; exists {
			commom = append(commom, letter)
		}
	}

	sort.Slice(commom, func(i, j int) bool {
		letterI := commom[i]
		letterJ := commom[j]
		return letterI < letterJ
	})

	return commom
}

func main() {
	s := "hello"
	letters := []rune{'h', 'a', 'e', 'i', 'o', 'u'}
	result := CommonCharacters(s, letters)
	fmt.Printf("%c, ", result) // Should print: e, h, o
}
