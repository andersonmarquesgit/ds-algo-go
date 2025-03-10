package main

import (
	"fmt"
)

type Deque struct {
	elements []string
}

func (d *Deque) EnqueueRight(element string) {
	d.elements = append(d.elements, element)
}

func (d *Deque) DequeueLeft() (string, bool) {
	if len(d.elements) == 0 {
		return "", false
	}
	element := d.elements[0]
	d.elements = d.elements[1:]
	return element, true
}

func (d *Deque) EnqueueLeft(element string) {
	d.elements = append([]string{element}, d.elements...)
}

func (d *Deque) DequeueRight() (string, bool) {
	if len(d.elements) == 0 {
		return "", false
	}
	element := d.elements[len(d.elements)-1]
	d.elements = d.elements[:len(d.elements)-1]
	return element, true
}

func main() {
	deque := &Deque{}

	deque.EnqueueRight("Orange")
	deque.EnqueueRight("Grapes")
	deque.EnqueueLeft("Apple")

	fmt.Println(*deque)

	if element, ok := deque.DequeueLeft(); ok {
		fmt.Println(element) // Expects "Apple"
	}

	if element, ok := deque.DequeueRight(); ok {
		fmt.Println(element) // Expects "Grapes"
	}
}
