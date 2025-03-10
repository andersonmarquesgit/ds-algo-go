package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Cosmo,is,an,incredible,technical,companion,with,very,strong,skills,in,Algorithms,and,Data,Structures,and,a,great,teacher,for,technical,interviews."
	words := strings.FieldsFunc(text, isDelimiter)

	// TODO: initialize the word count map
	wordCount := make(map[string]int)

	// TODO: count words and update the map
	for _, word := range words {
		wordCount[word]++
	}

	// TODO: print the word frequency
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func isDelimiter(r rune) bool {
	return r == ',' || r == '.'
}
