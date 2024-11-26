package main

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

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		newNode.next = l.head
		l.head.prev = &newNode
		l.tail = l.head
		l.head = &newNode
	}

	l.length++
}
