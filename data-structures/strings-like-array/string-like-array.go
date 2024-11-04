package main

import (
	"fmt"
	"strings"
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
	backwards := make([]string, len(s))
	totalItems := len(s) - 1
	for i := totalItems; i >= 0; i-- {
		backwards = append(backwards, string(s[i]))
	}

	justString := strings.Join(backwards, "")
	return justString
}

func main() {
	// Test the function
	s := "Hi My name is Anderson"
	reversed := reverseString(s)
	fmt.Println(reversed)

	reversed2 := reverseString2(s)
	fmt.Println(reversed2)

}
