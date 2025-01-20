package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isValidPalindrome(s string) bool {
	newString := replaceString(&s)
	left := 0
	right := len(newString) - 1

	for left <= right {
		if newString[left] != newString[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isValidPalindromeAlmost(s string) bool {
	newString := replaceString(&s)
	left := 0
	right := len(newString) - 1

	for left <= right {
		if newString[left] != newString[right] {
			return subValidPalindrome(newString, left+1, right) || subValidPalindrome(newString, left, right-1)
		}

		left++
		right--
	}

	return true
}

func subValidPalindrome(s string, left, right int) bool {
	for left <= right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func replaceString(s *string) string {
	// Compila o regex para encontrar tudo que não é alfanumérico
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	// Remove os caracteres que não são alfanuméricos
	return strings.ToLower(regex.ReplaceAllString(*s, ""))
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isValidPalindrome(s))

	s2 := "aabaa"
	fmt.Println(isValidPalindrome(s2))

	s3 := "aabbaa"
	fmt.Println(isValidPalindrome(s3))

	s4 := "abc"
	fmt.Println(isValidPalindrome(s4))

	s5 := "a"
	fmt.Println(isValidPalindrome(s5))

	fmt.Println("-- Almost Palindrome")

	s6 := "raceacar"
	fmt.Println(isValidPalindromeAlmost(s6))

	s7 := "abccdba"
	fmt.Println(isValidPalindromeAlmost(s7))
}
