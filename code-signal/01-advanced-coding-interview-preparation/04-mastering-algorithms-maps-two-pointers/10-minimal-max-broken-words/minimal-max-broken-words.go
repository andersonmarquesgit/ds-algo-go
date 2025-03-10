package main

import (
	"fmt"
	"strings"
)

func MinimalMaxBrokenWords(s string) (rune, int) {
	// Passo 1: Identificar palavras únicas
	wordsMap := make(map[string]bool)
	words := strings.Fields(s) // Divide a string em palavras baseadas em espaços

	for _, word := range words {
		wordsMap[word] = true // Usamos um map para garantir unicidade
	}

	// Passo 2: Contar quantas palavras cada caractere pode quebrar
	charImpact := make(map[rune]map[string]bool) // map[caractere] -> map[palavras afetadas]

	for word := range wordsMap { // Percorremos apenas palavras únicas
		for _, char := range word { // Percorremos os caracteres de cada palavra
			if _, exists := charImpact[char]; !exists {
				charImpact[char] = make(map[string]bool)
			}
			charImpact[char][word] = true
		}
	}

	// Passo 3: Encontrar o caractere que quebra mais palavras
	maxBrokenWords := 0
	var bestChar rune
	charFirstOccurrence := make(map[rune]int)

	// Percorrer a string original para garantir a primeira ocorrência
	for i, char := range s {
		if _, seen := charFirstOccurrence[char]; !seen {
			charFirstOccurrence[char] = i
		}
	}

	for char, affectedWords := range charImpact {
		brokenCount := len(affectedWords) // Número de palavras quebradas

		if brokenCount > maxBrokenWords || (brokenCount == maxBrokenWords && charFirstOccurrence[char] < charFirstOccurrence[bestChar]) {
			maxBrokenWords = brokenCount
			bestChar = char
		}
	}

	return bestChar, maxBrokenWords
}

func main() {
	resultChar, resultCount := MinimalMaxBrokenWords("Hello, world!")
	fmt.Printf("[%c, %d]\n", resultChar, resultCount)
	// Expected output: ['l', 2]
}
