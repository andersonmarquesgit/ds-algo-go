package main

import "fmt"

type MyDoubleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

func NewDoubleLinkedList() MyDoubleLinkedList {
	return MyDoubleLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *MyDoubleLinkedList) Append(value interface{}) {
	newNode := Node{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		// 1 | 2 | 3
		// 1 | 2 | 3 | New Node 4
		newNode.prev = l.tail
		l.tail.next = &newNode
		l.tail = &newNode
	}

	l.length++
}

func (l *MyDoubleLinkedList) Prepend(value interface{}) {
	newNode := Node{
		value: value,
		next:  nil,
		prev:  nil,
	}

	newNode.next = l.head
	l.head.prev = &newNode
	l.head = &newNode
	if l.tail == nil {
		l.tail = &newNode
	}

	l.length++
}

func (l *MyDoubleLinkedList) Insert(index int, value interface{}) {
	if index > 0 && index <= l.length {
		newNode := Node{
			value: value,
			prev:  nil,
			next:  nil,
		}

		currentPosition := 1
		currentNode := l.head

		if currentPosition == index && currentNode.next == nil {
			l.Append(value)
		} else {
			for currentPosition != index && currentNode.next != nil {
				currentNode = currentNode.next
				currentPosition++
			}

			if currentNode.next == nil {
				l.tail = &newNode
			}
			newNode.next = currentNode.next
			currentNode.next.prev = &newNode
			newNode.prev = currentNode
			currentNode.next = &newNode
			l.length++
		}
	}
}

func (l *MyDoubleLinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Print("| ", currentNode.value, " ")
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func (l *MyDoubleLinkedList) PrintReverse() {
	currentNode := l.tail
	for currentNode != nil {
		fmt.Print("| ", currentNode.value, " ")
		currentNode = currentNode.prev
	}
	fmt.Println("")
}

func main() {
	myDoubleLinkedList := NewDoubleLinkedList()
	myDoubleLinkedList.Append(1)
	myDoubleLinkedList.Append(2)
	myDoubleLinkedList.Append(3)
	myDoubleLinkedList.Append(4)
	myDoubleLinkedList.Insert(2, 5)
	myDoubleLinkedList.Print()        // 1 2 5 3 4
	myDoubleLinkedList.PrintReverse() // 4 3 5 2 1
	myDoubleLinkedList.Insert(1, 6)
	myDoubleLinkedList.Print()        // 1 6 2 5 3 4
	myDoubleLinkedList.PrintReverse() // 4 3 5 2 6 1
	myDoubleLinkedList.Prepend(7)
	myDoubleLinkedList.Print()        // 7 1 6 2 5 3 4
	myDoubleLinkedList.PrintReverse() // 4 3 5 2 6 1 7
}
