package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindromeValid(s string) bool {
	left, right := 0, len(s)-1
	sLower := strings.ToLower(s)
	// Aqui poderíamos pensar na remoção dos caracteres diferentes de alfanuméricos, mas a abordagem de andar com os ponteiros
	// também é possível, como a que fizemos abaixo validando caracter a caracter

	for left < right {
		for left < right && !isAlphanumeric(rune(sLower[left])) {
			left++
		}
		for left < right && !isAlphanumeric(rune(sLower[right])) {
			right--
		}
		if sLower[left] != sLower[right] {
			return false
		}

		left++
		right--
	}
	return true
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}

func main() {
	// Test Cases
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindromeValid(s), "|", s)

	s2 := "race a car"
	fmt.Println(isPalindromeValid(s2), "|", s2)

	s3 := "0P"
	fmt.Println(isPalindromeValid(s3), "|", s3)

	s4 := "0p"
	fmt.Println(isPalindromeValid(s4), "|", s4)

	s5 := ""
	fmt.Println(isPalindromeValid(s5), "|", s5)

	s6 := " "
	fmt.Println(isPalindromeValid(s6), "|", s6)

	s7 := "a"
	fmt.Println(isPalindromeValid(s7), "|", s7)

	s8 := "12.02.2021"
	fmt.Println(isPalindromeValid(s8), "|", s8)

	s9 := "race car"
	fmt.Println(isPalindromeValid(s9), "|", s9)
}
