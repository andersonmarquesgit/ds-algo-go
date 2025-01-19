package main

import "fmt"

/*
Time complexity: O(a + b)
Space complexity: O(a + b)
*/
func isSameString(s, t string) bool {
	resultS := buildString(s) // O(a)
	resultT := buildString(t) // O(b)

	if len(resultS) != len(resultT) {
		return false
	} else {
		for i := 0; i < len(resultS); i++ {
			// O(a) or O(b). Somado com a complexidade acima, temos: O(2a + b) ou O(a + 2b)
			// Então a nossa complexidade de tempo é O(a + b)
			if resultS[i] != resultT[i] {
				return false
			}
		}
	}

	return true
}

// O(n)
func buildString(str string) []byte {
	builtString := []byte{}
	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			builtString = append(builtString, str[i])
		} else {
			length := len(builtString)
			if length > 0 {
				builtString = builtString[:length-1]
			}
		}
	}

	return builtString
}

/*
Time Complexity: O(n + m)
Space Complexity: O(1)
*/
func isSameStringOptimized(s, t string) bool {
	p1, p2 := len(s)-1, len(t)-1

	for p1 >= 0 || p2 >= 0 {
		if (p1 >= 0 && s[p1] == '#') || (p2 >= 0 && t[p2] == '#') { // Se algum deles for um #, não queremos comparar, então resolvemos as strings
			if p1 >= 0 && s[p1] == '#' {
				backCount := 2
				for backCount > 0 {
					p1--
					backCount--
					if p1 >= 0 && s[p1] == '#' {
						backCount += 2
					}
				}
			}

			if p2 >= 0 && t[p2] == '#' {
				backCount := 2
				for backCount > 0 {
					p2--
					backCount--
					if p2 >= 0 && t[p2] == '#' {
						backCount += 2
					}
				}
			}
		} else {
			var charS, charT byte
			if p2 < 0 {
				charT = 0
			} else {
				charT = t[p2]
			}

			if p1 < 0 {
				charS = 0
			} else {
				charS = s[p1]
			}

			if charT != charS {
				return false
			} else {
				p1--
				p2--
			}
		}
	}

	return true
}

func main() {
	s := "ab#z"
	t := "az#z"
	fmt.Println(isSameString(s, t))

	s2 := "abc#d"
	t2 := "acc#d"
	fmt.Println(isSameString(s2, t2))

	s3 := "x#y#z#"
	t3 := "a#"
	fmt.Println(isSameString(s3, t3))

	s4 := "a###b"
	t4 := "b"
	fmt.Println(isSameString(s4, t4))

	s5 := "Ab#z"
	t5 := "ab#z"
	fmt.Println(isSameString(s5, t5))

	s6 := "Ab#z"
	t6 := "Az#z"
	fmt.Println(isSameString(s6, t6))

	s7 := "bxj##tw"
	t7 := "bxj###tw"

	fmt.Println("---- Using optimized function ----")
	fmt.Println(isSameStringOptimized(s, t))
	fmt.Println(isSameStringOptimized(s2, t2))
	fmt.Println(isSameStringOptimized(s3, t3))
	fmt.Println(isSameStringOptimized(s4, t4))
	fmt.Println(isSameStringOptimized(s5, t5))
	fmt.Println(isSameStringOptimized(s6, t6))
	fmt.Println(isSameStringOptimized(s7, t7))
}
