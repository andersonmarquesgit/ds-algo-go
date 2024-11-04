package main

import (
	"fmt"
)

// Create a function that reverses a string:
func reverseString(s string) string {
	// Convert string to rune slice
	r := []rune(s)

	// Reverse the slice
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// Return the reversed string
	return string(r)
}

func reverseString2(s string) string {
	// Criar um slice de runas de tamanho fixo igual ao comprimento da string
	r := []rune(s)
	backwards := make([]rune, len(s))

	// Preencher o slice ao contrário
	for i := 0; i < len(r); i++ {
		backwards[len(r)-1-i] = r[i]
	}

	// Converter o slice de volta para string
	return string(backwards)
}

func main() {
	// Test the function
	s := "Hi My name is Anderson"
	reversed := reverseString(s)
	fmt.Println(reversed)

	reversed2 := reverseString2(s)
	fmt.Println(reversed2)

}
