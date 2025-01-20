package main

/*
3. Longest Substring Without Repeating Characters
Medium
Topics
Companies
Hint
Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

import (
	"fmt"
	"strings"
)

/*
Time Complexity: O(n^2)
Space Complexity: O(n)
*/
func longestSubstring(str string) int {
	if len(str) <= 1 {
		return len(str)
	}

	longest := 0
	longestStr := []byte{}
	strLowerCase := strings.ToLower(str)

	for i := 0; i < len(strLowerCase); i++ {
		exist := false
		currentChar := strLowerCase[i]
		for _, char := range longestStr {
			if currentChar == char {
				exist = true
			}
		}

		if !exist {
			longestStr = append(longestStr, currentChar)
			if len(longestStr) > longest {
				longest = len(longestStr)
			}
		} else {
			longestStr = []byte{}
			longestStr = append(longestStr, currentChar)
		}
	}

	return longest
}

/*
Time Complexity: O(n)
Space Complexity: O(n)
*/
func longestSubstring2(str string) int {
	if len(str) <= 1 {
		return len(str)
	}

	longest := 0
	strMap := make(map[byte]bool)

	strLowerCase := strings.ToLower(str)
	for _, char := range strLowerCase {
		if _, ok := strMap[byte(char)]; !ok {
			strMap[byte(char)] = true
		} else {
			strMap = make(map[byte]bool)
			strMap[byte(char)] = true
		}

		if len(strMap) > longest {
			longest = len(strMap)
		}
	}

	return longest
}

/*
Time Complexity: O(n)
Space Complexity: O(n)
*/
func longestSubstringOptimized(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	longest := 0
	strMap := make(map[byte]int)
	leftWindow := 0
	var window string

	for rightWindow := 0; rightWindow < len(s); rightWindow++ {
		char := s[rightWindow]

		if charIndex, ok := strMap[char]; ok && charIndex >= leftWindow { // Character já foi visto e está dentro da janela
			leftWindow = charIndex + 1
		}

		strMap[char] = rightWindow

		// Nesse caso o uso de uma variável com potencial de aumentar a medida que a string aumenta, temos um
		// Space Complexity de O(n), caso precisássemos retornar a string formada
		window = s[leftWindow : rightWindow+1]
		if len(window) > longest {
			longest = len(window)
		}

		// Caso queiramos manter apenas o retorna do longest é possível um Space Complexity de O(1) e usarmos
		// Os limites da janela deslizante
		//if rightWindow-leftWindow+1 > longest {
		//	longest = rightWindow - leftWindow + 1
		//}

	}

	return longest
}

func main() {
	str1 := "abcbdca"
	fmt.Println(longestSubstring(str1))
	fmt.Println(longestSubstring2(str1))
	fmt.Println(longestSubstringOptimized(str1))

	str2 := "ccccc"
	fmt.Println(longestSubstring(str2))
	fmt.Println(longestSubstring2(str2))
	fmt.Println(longestSubstringOptimized(str2))

	str3 := ""
	fmt.Println(longestSubstring(str3))
	fmt.Println(longestSubstring2(str3))
	fmt.Println(longestSubstringOptimized(str3))

	str4 := "abcbda"
	fmt.Println(longestSubstring(str4))  // Função com erro para esse caso de teste, necessário janela deslizante
	fmt.Println(longestSubstring2(str4)) // Função com erro para esse caso de teste, necessário janela deslizante
	fmt.Println(longestSubstringOptimized(str4))

	str5 := "a"
	fmt.Println(longestSubstring(str5))
	fmt.Println(longestSubstring2(str5))
	fmt.Println(longestSubstringOptimized(str5))

}
