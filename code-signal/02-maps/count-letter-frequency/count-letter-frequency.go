package main

import (
	"fmt"
	"unicode"
)

// Function to count the frequency of each letter in a given sentence
func CountLetterFrequency(sentence string) map[rune]int {
	letterCount := make(map[rune]int)
	for _, letter := range sentence {
		// TODO: If the character is a letter, update its count in the map
		if unicode.IsLetter(letter) {
			letterCount[unicode.ToLower(letter)]++
		}
		// or add it with a count of 1 if it's not already there
	}
	return letterCount
}

func main() {
	sentence := "Once upon a time in a faraway library"
	// TODO: Call the function with the sentence variable and print the result
	result := CountLetterFrequency(sentence)
	for key, value := range result {
		fmt.Printf("%s: %d\n", string(key), value)
	}
}
