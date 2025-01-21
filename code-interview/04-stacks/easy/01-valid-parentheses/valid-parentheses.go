package main

import "fmt"

/*
Time Complexity: O(n)
Space Complexity: O(n)
*/
func validParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}

	brackets := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	stack := []byte{} // O(n)

	for i := 0; i < len(s); i++ { // O(n)
		currentBracket := s[i]
		if _, found := brackets[currentBracket]; found { // Significa que recebemos um bracket esquerdo
			stack = append(stack, currentBracket)
		} else { // Significa que recebemos um bracket direito
			if len(stack) == 0 {
				return false // Nenhum esquerdo correspondente
			}
			leftBracket := stack[len(stack)-1]
			rightBracket := brackets[leftBracket]

			if rightBracket != currentBracket {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

func main() {
	s := "{([])}"
	fmt.Println(validParentheses(s))

	s2 := "{()[]}"
	fmt.Println(validParentheses(s2))

	s3 := ""
	fmt.Println(validParentheses(s3))

	s4 := "{([)]}"
	fmt.Println(validParentheses(s4))

	s5 := "{([])([)]}"
	fmt.Println(validParentheses(s5))

	s6 := "})]"
	fmt.Println(validParentheses(s6))

	s7 := "}"
	fmt.Println(validParentheses(s7))
}
