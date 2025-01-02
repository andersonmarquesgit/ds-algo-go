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

func main() {
	fmt.Println(FirstReverse("I Love Code"))
}
