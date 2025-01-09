package main

import (
	"fmt"
	"strconv"
)

/*
We have the input string: aaaabbccc
And the output is: 4a2b3c

Clarify:
- If we have a same character on the end array, for example, the character 'a' again: aaaabbccca
The output is 4a2b3c1a

- If we have input abcd
The output is 1a1b1c1d

*/

func encode(str string) string { // O(n)
	if len(str) == 0 {
		return ""
	}

	strEncode := ""
	countChar := 0
	currentChar := str[0]

	for i := 0; i < len(str); i++ {
		if currentChar == str[i] {
			countChar++
		} else {
			strEncode += strconv.Itoa(countChar) + string(currentChar)
			currentChar = str[i]
			countChar = 1
		}
	}

	strEncode += strconv.Itoa(countChar) + string(currentChar) // Add the last character
	return strEncode
}

func main() {
	str := "aaaabbccc"
	fmt.Println(encode(str))

	str2 := "abcd"
	fmt.Println(encode(str2))

	str3 := "aaaabbccca"
	fmt.Println(encode(str3))

	str4 := "a"
	fmt.Println(encode(str4))
}
