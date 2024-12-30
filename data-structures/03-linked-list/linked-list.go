package main

import "fmt"

type LinkedList struct {
	head    *Node
	length  int
	tail    *Node
	nodeMap map[int]*Node // Mapear índices para nós e ter acesso O(1)
	offset  int           // Deslocamento lógico para a posição 0
}

type Node struct {
	value interface{}
	next  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:    nil,
		length:  0,
		tail:    nil,
		offset:  0,
		nodeMap: make(map[int]*Node),
	}
}

func (l *LinkedList) GetNode(position int) *Node {
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
		l.nodeMap[0] = &newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil { // O(n)
			currentNode = currentNode.next
		}
		currentNode.next = &newNode
		l.nodeMap[l.length] = &newNode
	}

	l.tail = &newNode
	l.length++
}

// generate a append2 method with O(1) complexity
func (l *LinkedList) Append2(value interface{}) {
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

func (l *LinkedList) Prepend(value interface{}) {
	newNode := Node{
		value: value,
		next:  l.head,
	}

	l.head = &newNode
	realIndex := 0
	if l.offset == 0 {
		realIndex = 0
	} else {
		realIndex = l.offset
	}
	l.nodeMap[realIndex] = &newNode
	l.offset--
	l.length++

	// 35 | 101
	// 0  |  1 => nodeMap and offset
	// 100 | 35 | 101
	// -1  |  0 |  1 => nodeMap and offset
	// 100 | 35 | [55] | 101 => 55 with insert method in position 2
	// -1  |  0 |   1  |  2 => nodeMap and offset

	if l.tail == nil {
		l.tail = &newNode
	}
}

func (l *LinkedList) Insert(index int, value interface{}) {
	currentPosition := 1
	if index > 0 && index <= l.length {
		newNode := Node{
			value: value,
			next:  nil,
		}

		currentNode := l.head
		if currentPosition == index && currentNode.next == nil {
			currentNode.next = &newNode
			l.tail = &newNode
			l.nodeMap[l.length] = &newNode
			l.length++
		} else {
			for currentPosition != index && currentNode.next != nil { // O(n)
				currentNode = currentNode.next
				currentPosition++
			}

			if currentNode.next == nil {
				l.tail = &newNode
			}
			newNode.next = currentNode.next
			currentNode.next = &newNode
			l.nodeMap[l.length] = &newNode
			l.length++
		}
	}
}

// generate a insert2 method with O(1) complexity
func (l *LinkedList) Insert2(index int, value interface{}) {
	if index < 0 || index > l.length {
		panic("Invalid index")
	}

	if index == 0 {
		l.Prepend(value)
		return
	}

	if index == l.length {
		l.Append2(value)
		return
	}

	newNode := Node{
		value: value,
		next:  nil,
	}

	// Calcular o índice real no nodeMap
	//realIndex := index - 1 + l.offset

	realIndex := 0
	if l.offset == 0 {
		realIndex = index - 1
	} else {
		realIndex = index + l.offset
	}

	// Obter o nó anterior
	prevNode := l.nodeMap[realIndex]

	newNode.next = prevNode.next
	prevNode.next = &newNode
	//l.nodeMap[index+l.offset] = &newNode
	l.nodeMap[realIndex] = &newNode
	l.length++
}

func (l *LinkedList) Remove(index int) {
	if index < 0 || index >= l.length {
		panic("Invalid index")
	}

	if index == 0 {
		l.head = l.head.next
		l.length--
		return
	}

	currentPosition := 1
	currentNode := l.head
	for currentPosition != index && currentNode.next != nil {
		currentNode = currentNode.next
		currentPosition++
	}

	currentNode.next = currentNode.next.next
	l.length--
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
	myLinkedList.Prepend(35)
	myLinkedList.Print()
	fmt.Println("Insert With Prepend: ", 35)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	myLinkedList.Insert2(1, 101)
	myLinkedList.Print()
	fmt.Println("Insert With Insert: ", 101, " in index ", 1)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	myLinkedList.Insert2(0, 100)
	myLinkedList.Print()
	fmt.Println("Insert With Insert: ", 100, " in index ", 0)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	myLinkedList.Insert2(2, 55)
	myLinkedList.Print()
	fmt.Println("Insert With Insert: ", 55, " in index ", 2)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	myLinkedList.Append(888)
	myLinkedList.Print()
	fmt.Println("Insert With Append: ", 888)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	myLinkedList.Prepend(111)
	myLinkedList.Print()
	fmt.Println("Insert With Prepend: ", 111)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	myLinkedList.Insert2(2, 222)
	myLinkedList.Print()
	fmt.Println("Insert With Insert: ", 222, " in index ", 2)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	myLinkedList.Remove(2)
	myLinkedList.Print()
	fmt.Println("Deletion index: ", 2)
	fmt.Println("Offset: ", myLinkedList.offset)
	fmt.Println("Length: ", myLinkedList.length)
	fmt.Println("Head: ", myLinkedList.head.value)
	fmt.Println("Tail: ", myLinkedList.tail.value)
	fmt.Println("------")

	//myLinkedList.Append2(5)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Append2(10)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Append2(15)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Prepend(35)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Insert(3, 55)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Insert(5, 11)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Insert(1, 101)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Insert2(0, 100)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
	//myLinkedList.Insert(3, 333)
	//myLinkedList.Print()
	//fmt.Println("Head: ", myLinkedList.head.value)
	//fmt.Println("Next: ", myLinkedList.head.next.value)
	//fmt.Println("Tail: ", *myLinkedList.tail)
	//fmt.Println("Length", myLinkedList.length)
	//fmt.Println("------")
}
