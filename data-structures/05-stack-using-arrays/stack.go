package main

import "fmt"

/** create stack using arrays */

type Stack struct {
	nodes []interface{}
}

func NewStack() *Stack {
	return &Stack{
		nodes: []interface{}{},
	}
}

func (s *Stack) peek() interface{} {
	if s.isEmpty() {
		return nil
	}

	return s.nodes[len(s.nodes)-1]
}

func (s *Stack) push(value interface{}) {
	s.nodes = append(s.nodes, value)
}

func (s *Stack) pop() {
	if s.isEmpty() {
		return
	}

	s.nodes = s.nodes[:len(s.nodes)-1] // O(1) - slice operation to remove the last element of the array (stack)

}

func (s *Stack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *Stack) print() {
	for i := len(s.nodes) - 1; i >= 0; i-- {
		fmt.Println(s.nodes[i])
	}
}

func main() {
	/*
		Like a history navigation on browser
		- discord
		- udemy
		- google
	*/
	myStack := NewStack()
	myStack.push("discord.com")
	myStack.push("udemy.com")
	myStack.push("google.com")
	myStack.print()

	fmt.Println("---------")
	myStack.pop()
	myStack.print()

	fmt.Println("---------")
	myStack.push("amazon.com")
	fmt.Println(myStack.peek())

	fmt.Println("---------")
	myStack.print()
}
