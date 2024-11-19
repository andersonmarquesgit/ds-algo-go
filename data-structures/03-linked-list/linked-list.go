package main

import "fmt"

type LinkedList struct {
	head   *Node
	length int
	tail   *Node
}

type Node struct {
	value interface{}
	next  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Append(value interface{}) {
	newNode := Node{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = &newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &newNode
	}

	l.tail = &newNode
	l.length++
}

func (l *LinkedList) Prepend(value interface{}) {
	newNode := Node{
		value: value,
		next:  l.head,
	}

	l.head = &newNode
	l.length++
}

func (l *LinkedList) Insert(position int, value interface{}) {
	currentPosition := 1
	if position > 0 && position <= l.length {
		newNode := Node{
			value: value,
			next:  nil,
		}

		currentNode := l.head
		if currentPosition == position && currentNode.next == nil {
			currentNode.next = &newNode
			l.tail = &newNode
			l.length++
		} else {
			for currentPosition != position && currentNode.next != nil {
				currentNode = currentNode.next
				currentPosition++
			}
			newNode.next = currentNode.next
			currentNode.next = &newNode
			l.length++
		}
	}
}

func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Print("| ", currentNode.value, " ")
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func main() {
	myLinkedList := NewLinkedList()
	myLinkedList.Append(5)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Next: ", myLinkedList.head.next)
	fmt.Println("Tail: ", *myLinkedList.tail)
	fmt.Println("Length", myLinkedList.length)
	fmt.Println("------")
	myLinkedList.Append(10)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Next: ", myLinkedList.head.next.value)
	fmt.Println("Tail: ", *myLinkedList.tail)
	fmt.Println("Length", myLinkedList.length)
	fmt.Println("------")
	myLinkedList.Append(15)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Next: ", myLinkedList.head.next.value)
	fmt.Println("Tail: ", *myLinkedList.tail)
	fmt.Println("Length", myLinkedList.length)
	fmt.Println("------")
	myLinkedList.Prepend(35)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Next: ", myLinkedList.head.next.value)
	fmt.Println("Tail: ", *myLinkedList.tail)
	fmt.Println("Length", myLinkedList.length)
	fmt.Println("------")
	myLinkedList.Insert(3, 55)
	myLinkedList.Print()
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Next: ", myLinkedList.head.next.value)
	fmt.Println("Tail: ", *myLinkedList.tail)
	fmt.Println("Length", myLinkedList.length)
	fmt.Println("------")
	myLinkedList.Insert(5, 11)
	myLinkedList.Print()
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Next: ", myLinkedList.head.next.value)
	fmt.Println("Tail: ", *myLinkedList.tail)
	fmt.Println("Length", myLinkedList.length)
	fmt.Println("------")

}
