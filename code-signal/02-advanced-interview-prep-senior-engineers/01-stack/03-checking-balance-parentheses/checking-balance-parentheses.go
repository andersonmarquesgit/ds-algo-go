package main

import (
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func IsParenBalanced(parenString string) bool {
	var stack Stack
	openingParen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, paren := range parenString {
		if paren == '(' || paren == '[' || paren == '{' {
			// We met an opening parenthesis, just putting it on the stack
			stack.Push(paren)
		} else if closing, exists := openingParen[paren]; exists {
			// We met a closing parenthesis
			if len(stack.items) == 0 || stack.items[len(stack.items)-1] != closing {
				return false
			}
			stack.Pop()
		}
	}

	return len(stack.items) == 0
}

func main() {
	fmt.Println(IsParenBalanced("(())"))  // Outputs: true
	fmt.Println(IsParenBalanced("({[)}")) // Outputs: false
}
