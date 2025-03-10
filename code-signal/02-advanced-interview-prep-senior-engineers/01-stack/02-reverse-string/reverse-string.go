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

func Reverse(input string) string {
	var stack Stack

	for _, c := range input {
		stack.Push(c)
	}

	var reversedString []rune
	for len(stack.items) > 0 {
		reversedString = append(reversedString, stack.Pop())
	}

	return string(reversedString)
}

func main() {
	fmt.Println(Reverse("HELLO")) // Outputs: OLLEH
}
