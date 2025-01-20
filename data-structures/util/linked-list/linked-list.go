package linked_list

import "fmt"

type LinkedList struct {
	Head    *Node
	length  int
	Tail    *Node
	nodeMap map[int]*Node // Mapear índices para nós e ter acesso O(1)
	offset  int           // Deslocamento lógico para a posição 0
}

type Node struct {
	Value interface{}
	Next  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:    nil,
		length:  0,
		Tail:    nil,
		offset:  0,
		nodeMap: make(map[int]*Node),
	}
}

func (l *LinkedList) GetNode(position int) *Node {
	realIndex := position + l.offset
	return l.nodeMap[realIndex]
}

// Time complexity: O(1), Node sempre inserido na Tail
func (l *LinkedList) Append(value interface{}) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
		l.nodeMap[0] = &newNode
		return
	}

	l.Tail.Next = &newNode
	l.Tail = &newNode
	l.nodeMap[l.length] = &newNode
	l.length++
}

// Time complexity: O(1), Node sempre inserido no Head
func (l *LinkedList) Prepend(value interface{}) {
	newNode := Node{
		Value: value,
		Next:  l.Head,
	}

	l.Head = &newNode
	realIndex := 0
	if l.offset == 0 {
		realIndex = 0
	} else {
		realIndex = l.offset
	}
	l.nodeMap[realIndex] = &newNode
	l.offset--
	l.length++

	if l.Tail == nil {
		l.Tail = &newNode
	}
}

func (l *LinkedList) Insert(index int, value interface{}) {
	currentPosition := 1
	if index > 0 && index <= l.length {
		newNode := Node{
			Value: value,
			Next:  nil,
		}

		currentNode := l.Head
		if currentPosition == index && currentNode.Next == nil {
			currentNode.Next = &newNode
			l.Tail = &newNode
			l.nodeMap[l.length] = &newNode
			l.length++
		} else {
			for currentPosition != index && currentNode.Next != nil { // O(n)
				currentNode = currentNode.Next
				currentPosition++
			}

			if currentNode.Next == nil {
				l.Tail = &newNode
			}
			newNode.Next = currentNode.Next
			currentNode.Next = &newNode
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
		l.Append(value)
		return
	}

	newNode := Node{
		Value: value,
		Next:  nil,
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

	newNode.Next = prevNode.Next
	prevNode.Next = &newNode
	//l.nodeMap[index+l.offset] = &newNode
	l.nodeMap[realIndex] = &newNode
	l.length++
}

func (l *LinkedList) Remove(index int) {
	if index < 0 || index >= l.length {
		panic("Invalid index")
	}

	if index == 0 {
		l.Head = l.Head.Next
		l.length--
		return
	}

	currentPosition := 1
	currentNode := l.Head
	for currentPosition != index && currentNode.Next != nil {
		currentNode = currentNode.Next
		currentPosition++
	}

	currentNode.Next = currentNode.Next.Next
	l.length--
}

func (l *LinkedList) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Print("| ", currentNode.Value, " ")
		currentNode = currentNode.Next
	}
	fmt.Println("")
}

func (l *LinkedList) PrintByNode(head *Node) {
	currentNode := head
	for currentNode != nil {
		fmt.Print("| ", currentNode.Value, " ")
		currentNode = currentNode.Next
	}
	fmt.Println("")
}
