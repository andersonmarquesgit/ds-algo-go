package main

import "fmt"

/*
100. Same Tree
Easy
Topics
Companies
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:

Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:

Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:

Input: p = [1,2,1], q = [1,1,2]
Output: false

Constraints:

The number of nodes in both trees is in the range [0, 100].
-104 <= TreeNode.val <= 104
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // Base case - If both nodes are nil,
		return true
	}

	if p == nil || q == nil { // Base case - If one of the nodes is nil,
		return false
	}

	if p.Value != q.Value { // Base case
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) // Recursive case - Check the left and right subtrees
}

func main() {
	myBST := NewBinarySearchTree()
	myBST.Insert(9)
	myBST.Insert(4)
	myBST.Insert(6)

	myBST2 := NewBinarySearchTree()
	myBST2.Insert(9)
	myBST2.Insert(4)
	myBST2.Insert(6)

	fmt.Println(isSameTree(myBST.root, myBST2.root))
}

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
