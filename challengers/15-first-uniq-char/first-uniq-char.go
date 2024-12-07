/*
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Example 1:

Input: s = "leetcode"

Output: 0

Explanation:

The character 'l' at index 0 is the first character that does not occur at any other index.

Example 2:

Input: s = "loveleetcode"

Output: 2

Example 3:

Input: s = "aabb"

Output: -1

Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.
*/
package main

import "fmt"

func firstUniqChar(s string) int {
	runeSlice := []rune(s)
	set := make(map[int]int) // Map for the index char

	// Create a map to store the index of each char in the string
	for i, v := range runeSlice {
		if _, ok := set[int(v)]; ok {
			set[int(v)] = -1
		} else {
			set[int(v)] = i
		}
	}

	// Iterate over the map and return the first char that has an index different from -1
	for _, v := range runeSlice {
		if set[int(v)] != -1 {
			return set[int(v)]
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
