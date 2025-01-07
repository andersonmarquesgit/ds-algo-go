package main

import (
	"fmt"
	"strings"
)

type BinarySearchTree struct {
	root *TreeNode
}

type TreeNode struct {
	Value int       `json:"value"`
	Left  *TreeNode `json:"left,omitempty"`
	Right *TreeNode `json:"right,omitempty"`
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func NewNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (tree *BinarySearchTree) Insert(value int) {
	newNode := NewNode(value)
	if tree.root == nil {
		tree.root = newNode
		return
	} else {
		insertRecursion(newNode, tree.root) // Poderíamos ter usado um loop infinito aqui, mas optamos por usar recursão
	}
}

func insertRecursion(newNode, currentNode *TreeNode) {
	if currentNode.Value > newNode.Value { // go to Left
		// Verify Left node
		if currentNode.Left == nil {
			currentNode.Left = newNode
		} else {
			insertRecursion(newNode, currentNode.Left)
		}
	} else { // go to Right
		// Verify Right node
		if currentNode.Right == nil {
			currentNode.Right = newNode
		} else {
			insertRecursion(newNode, currentNode.Right)
		}
	}
}

func levelOrder(queue []*TreeNode, list []int) {
	if len(queue) == 0 {
		output := strings.Trim(fmt.Sprint(list), "[]")
		fmt.Println(output)
		return
	}

	currentNode := queue[0]
	queue = queue[1:]

	list = append(list, currentNode.Value)

	if currentNode.Left != nil {
		queue = append(queue, currentNode.Left)
	}
	if currentNode.Right != nil {
		queue = append(queue, currentNode.Right)
	}

	levelOrder(queue, list)
}

func main() {
	var n int
	fmt.Scan(&n)
	myBST := NewBinarySearchTree()

	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		myBST.Insert(value)
	}

	levelOrder([]*TreeNode{myBST.root}, []int{})
}
