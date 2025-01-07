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

func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func (l *LinkedList) Prepend(value interface{}) {
	newNode := Node{
		value: value,
		next:  l.head,
	}

	l.head = &newNode
	realIndex := int32(0)
	if l.offset == 0 {
		realIndex = 0
	} else {
		realIndex = l.offset
	}
	l.nodeMap[realIndex] = &newNode
	l.offset--
	l.length++

	if l.tail == nil {
		l.tail = &newNode
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	myLinkedList := NewLinkedList()

	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		myLinkedList.Prepend(value)
	}

	myLinkedList.Print()
}
