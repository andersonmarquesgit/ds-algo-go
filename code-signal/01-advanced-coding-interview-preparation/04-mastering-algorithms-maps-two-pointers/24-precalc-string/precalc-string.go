package main

import (
	"fmt"
)

func Solution(S string, Q []struct{ c1, c2 rune }) []int {
	// TODO: Implement the solution
	resMap := preprocessString(S)
	result := make([]int, 0, len(Q))

	for _, q := range Q {
		key := string(q.c1) + string(q.c2)
		if val, found := resMap[key]; found {
			result = append(result, val)
		} else {
			result = append(result, 0)
		}
	}

	return result
}

func preprocessString(s string) map[string]int {
	resMap := make(map[string]int)

	for char1 := 'a'; char1 <= 'z'; char1++ {
		for char2 := char1 + 1; char2 <= 'z'; char2++ {
			key := string(char1) + string(char2)
			//fmt.Println("Processing pair:", key)
			resMap[key] = getMaxSequenceWithoutChars(s, char1, char2)
		}
	}

	return resMap
}

func getMaxSequenceWithoutChars(s string, c1, c2 rune) int {
	maxLength, currentLength := 0, 0

	for _, c := range s {
		if c == c1 || c == c2 {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 0
		} else {
			currentLength++
		}
	}

	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}

func main() {
	result1 := Solution("abcccacba", []struct{ c1, c2 rune }{{'a', 'b'}, {'b', 'c'}})
	fmt.Println(result1) // Example output: [3, 1]

	result2 := Solution("intelliaiassistant", []struct{ c1, c2 rune }{{'a', 'i'}, {'n', 't'}})
	fmt.Println(result2) // Example output: [5, 11]
}
