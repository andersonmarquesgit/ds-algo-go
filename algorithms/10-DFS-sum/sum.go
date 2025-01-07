package main

import "fmt"

func sumTree(tree *BinarySearchTree) int {
	return sum(tree.root)
}

func sum(node *Node) int {
	if node == nil { // Base case
		return 0
	}

	left := sum(node.Left)
	right := sum(node.Right)
	return left + right + node.Value
}

func main() {
	myBST := NewBinarySearchTree()
	myBST.Insert(1)
	myBST.Insert(2)
	myBST.Insert(3)
	myBST.Insert(4)
	myBST.Insert(5)
	myBST.Insert(6)

	fmt.Println("Total sum tree:", sumTree(myBST))
}

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	Value int   `json:"value"`
	Left  *Node `json:"left,omitempty"`
	Right *Node `json:"right,omitempty"`
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func NewNode(value int) *Node {
	return &Node{
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

func insertRecursion(newNode, currentNode *Node) {
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
