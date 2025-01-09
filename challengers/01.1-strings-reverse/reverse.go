package main

import "fmt"

func FirstReverse(str string) string {
	strRune := []rune(str)
	strLength := len(strRune)
	strReverse := make([]rune, strLength)

	for i := 0; i < strLength; i++ {
		strReverse[strLength-1-i] = strRune[i]
	}

	return string(strReverse)
}

func FirstReverse2(str string) string {
	strLength := len(str)
	strReverse := make([]byte, strLength)

	for i, j := 0, strLength-1; i < strLength; i, j = i+1, j-1 {
		strReverse[j] = str[i]
	}

	return string(strReverse)
}

func main() {
	fmt.Println(FirstReverse("I Love Code"))

	fmt.Println(FirstReverse2("I Love Code"))
}
