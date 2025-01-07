package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head    *Node
	length  int32
	tail    *Node
	nodeMap map[int32]*Node
	offset  int32
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:    nil,
		length:  0,
		tail:    nil,
		offset:  0,
		nodeMap: make(map[int32]*Node),
	}
}

func (l *LinkedList) GetNode(position int32) *Node {
	realIndex := position + l.offset
	return l.nodeMap[realIndex]
}

func (l *LinkedList) Append(value interface{}) {
	newNode := Node{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		l.nodeMap[0] = &newNode
		return
	}

	l.tail.next = &newNode
	l.tail = &newNode
	l.nodeMap[l.length] = &newNode
	l.length++
}

func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func main() {
	var n int
	fmt.Scan(&n)
	myLinkedList := NewLinkedList()

	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		myLinkedList.Append(value)
	}

	myLinkedList.Print()
}
