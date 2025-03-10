package main

import (
	"fmt"
)

func LongestSubstringWithKDistinctCharacters(s string, K int) int {
	// TODO: We need find the longest substring with K distinct characteres
	// For example:
	// "acaabcc" and K = 2 ---> acaa is the longest with 4 characteres
	// "acaabcc" and K = 1 ---> aa and cc is the longest with 2 characteres
	// "aaaaaaa" and K = 1 ---> aaaaaaa is the longest with 7 characteres
	// TODO: We need a map for count characteres posible -> For example: "acaabcc" and K = 2 -> a:true, c:true, b:true (excedd K) so clean the map
	// TODO: We can use sliding window, len(window) is greater than longest, so longest = len(window) => window = s[leftWindow:rightWindow+1]
	// TODO: If clean the map for count characteres, movement the leftWindor for current position

	if len(s) <= 1 {
		if K == 0 {
			return 0
		}
		return len(s)
	}

	longest := 0
	charCount := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		// Se houver mais de K caracteres distintos, movemos o `left`
		for len(charCount) > K {
			charCount[s[left]]--
			if charCount[s[left]] == 0 { // Remove do mapa se não houver mais ocorrências
				delete(charCount, s[left])
			}
			left++
		}

		longest = max(longest, right-left+1)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(LongestSubstringWithKDistinctCharacters("acaabcc", 2)) // Expected output: 4
	fmt.Println(LongestSubstringWithKDistinctCharacters("abaccc", 2))  // Expected output: 4
	fmt.Println(LongestSubstringWithKDistinctCharacters("aa", 1))      // Expected output: 2
}
