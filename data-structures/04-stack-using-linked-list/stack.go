package main

import "fmt"

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

type Node struct {
	value interface{}
	next  *Node
}

func NewStack() Stack {
	return Stack{
		top:    nil,
		bottom: nil,
		length: 0,
	}
}

func NewNode(value interface{}) Node {
	return Node{
		value: value,
		next:  nil,
	}
}

func (s *Stack) peek() *Node {
	return s.top
}

func (s *Stack) push(value interface{}) {
	newNode := NewNode(value)
	if s.isEmpty() {
		s.top = &newNode
		s.bottom = &newNode
		s.length++
		return
	}

	holdingPointer := s.top
	s.top = &newNode
	s.top.next = holdingPointer
	s.length++
}

func (s *Stack) pop() {
	if s.isEmpty() {
		s.bottom = nil
		return
	}

	s.top = s.top.next
	s.length--

	// 1 | 2 | 3
	// top = 1
	// bottom = 3
	// pop is remove 1, top position

	// 2 | 3
	// top = 2 => 1.next
	// bottom = 3
}

func (s *Stack) isEmpty() bool {
	if s.length == 0 {
		return true
	}

	return false
}

func (s *Stack) print() {
	if s.length == 0 {
		fmt.Println("Is empty!")
		return
	}

	currentNode := s.top
	for currentNode.next != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println(currentNode.value)
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
	fmt.Println(myStack.peek().value)

	fmt.Println("---------")
	myStack.print()

	fmt.Println("---------")
	fmt.Println("Pop", myStack.peek().value)
	myStack.pop()
	myStack.print()
	fmt.Println("---------")
	fmt.Println("Pop", myStack.peek().value)
	myStack.pop()
	myStack.print()
	fmt.Println("---------")
	fmt.Println("Pop", myStack.peek().value)
	myStack.pop()
	myStack.print()
}
