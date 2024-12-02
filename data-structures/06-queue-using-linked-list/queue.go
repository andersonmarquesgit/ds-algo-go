package main

import "fmt"

type Queue struct {
	first  *Node
	last   *Node
	length int
}

type Node struct {
	value interface{}
	next  *Node
}

func NewQueue() *Queue {
	return &Queue{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func (q *Queue) Peek() interface{} {
	if q.isEmpty() {
		return nil
	}

	return q.first.value
}

func (q *Queue) Enqueue(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	if q.isEmpty() {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}

	q.length++
}

func (q *Queue) Dequeue() {
	if q.isEmpty() {
		return
	}

	if q.first == q.last {
		q.last = nil
	}

	q.first = q.first.next
	q.length--
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}

func (q *Queue) Print() {
	currentNode := q.first
	for currentNode != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
}

func main() {
	myQueue := NewQueue()
	myQueue.Enqueue("discord")
	myQueue.Enqueue("udemy")
	myQueue.Enqueue("google")
	myQueue.Print()

	println("---------")
	myQueue.Dequeue()
	myQueue.Print()

	println("---------")
	myQueue.Dequeue()
	myQueue.Print()
}
