package main

import "fmt"

func reverseStringRecursion(str string) string {
	if len(str) == 0 {
		return ""
	}

	return reverseStringRecursion(str[1:]) + string(str[0])
	// This is the recursion with the string slicing and concatenation
	// The base case is when the string is empty, it returns an empty string
	// The recursion is the concatenation of the last character of the string with the rest of the string
	// The recursion will keep going until the string is empty
	// For example, if the string is "hello", the recursion will be:
	// reverseStringRecursion("hello") -> reverseStringRecursion("ello") + "h"
	// reverseStringRecursion("ello") -> reverseStringRecursion("llo") + "e"
	// reverseStringRecursion("llo") -> reverseStringRecursion("lo") + "l"
	// reverseStringRecursion("lo") -> reverseStringRecursion("o") + "l"
	// reverseStringRecursion("o") -> reverseStringRecursion("") + "o"
	// reverseStringRecursion("") -> ""
	// The recursion will return "olleh"
}

func reverseStringIterative(str string) string {
	var reversed string

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	return reversed
	// This is the iterative way to reverse a string
	// The reversed string is empty
	// The loop will iterate over the string from the last character to the first character
	// The loop will concatenate the last character of the string with the reversed string
	// The loop will keep going until the index is 0
	// For example, if the string is "hello", the loop will be:
	// reversed = "" + "o"
	// reversed = "o" + "l"
	// reversed = "ol" + "l"
	// reversed = "oll" + "e"
	// reversed = "olle" + "h"
	// The loop will return "olleh"
}

func reverseString(s []byte) {
	reverseStringHelper(s, 0, len(s)-1)
}

func reverseStringHelper(s []byte, left, right int) {
	if left >= right {
		return
	}

	s[left], s[right] = s[right], s[left]
	reverseStringHelper(s, left+1, right-1)
}

func main() {
	fmt.Println(reverseStringRecursion("hello"))  // olleh
	fmt.Println(reverseStringRecursion("world"))  // dlrow
	fmt.Println(reverseStringRecursion("golang")) // gnalog

	fmt.Println(reverseStringIterative("hello"))  // olleh
	fmt.Println(reverseStringIterative("world"))  // dlrow
	fmt.Println(reverseStringIterative("golang")) // gnalog

	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s)) // olleh
}
