package main

import (
	"fmt"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}
	stack.Push("Steve")
	stack.Push("Sam")

	fmt.Println(stack.Peek()) // Outputs: Sam

	fmt.Println(stack.IsEmpty()) // Outputs: false
	stack.Pop()                  // Remove 'Sam'
	stack.Pop()                  // Remove 'Steve'
	fmt.Println(stack.IsEmpty()) // Outputs: true
}
