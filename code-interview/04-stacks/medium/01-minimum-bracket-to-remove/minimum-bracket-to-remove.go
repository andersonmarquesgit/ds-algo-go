package main

import (
	"fmt"
	"sort"
)

/*
Time Complexity: O(n)
Space Complexity: O(n + m log m) ==> O(n)
*/
func minimumRemoveToMakeValid(s string) string {
	stack := []int{}       // space complexity O(n)
	newString := []byte(s) // space complexity O(n)

	for i := 0; i < len(newString); i++ { // time complexity O(n)
		// We need valid if the char is left or right
		currentChar := newString[i]

		if currentChar == '(' {
			stack = append(stack, i)
		} else if currentChar == ')' {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				index := stack[len(stack)-1]
				if newString[index] == '(' {
					stack = stack[:len(stack)-1] // remove the last open bracket, because is open bracket
				} else {
					stack = append(stack, i)
				}
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(stack))) // space complexity O(m log m)

	// Remove os caracteres nos índices especificados
	for _, index := range stack {
		newString = append(newString[:index], newString[index+1:]...)
	}

	return string(newString)
}

func main() {
	s1 := "a)bc(d)"
	fmt.Println(minimumRemoveToMakeValid(s1))

	s2 := "(ab(c)d"
	fmt.Println(minimumRemoveToMakeValid(s2))

	s3 := "))(("
	fmt.Println(minimumRemoveToMakeValid(s3))

}
