package main

import (
	"encoding/json"
	"fmt"
)

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

func (tree *BinarySearchTree) Lookup(value int) bool {
	if tree.root == nil {
		return false
	}

	currentNode := tree.root
	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return true
		}
	}
	return false
}

func (tree *BinarySearchTree) Delete(value int) {
	if tree.root == nil {
		return
	}

	// Find the node to delete
	var currentNode, parentNode *Node
	currentNode = tree.root
	for currentNode != nil {
		if value < currentNode.Value { // Go to Left
			parentNode = currentNode
			currentNode = currentNode.Left
		} else if value > currentNode.Value { // Go to Right
			parentNode = currentNode
			currentNode = currentNode.Right
		} else { // Found the node to delete
			// We found the node to delete
			// Option 1: No right child
			if currentNode.Right == nil {
				if parentNode == nil { // This case is when we are deleting the root node
					tree.root = currentNode.Left
				} else {
					if currentNode.Value < parentNode.Value { // Node to delete is on the Left of parent node
						parentNode.Left = currentNode.Left
					} else if currentNode.Value > parentNode.Value { // Node to delete is on the Right of parent node
						parentNode.Right = currentNode.Left
					}
				}
			} else if currentNode.Right.Left == nil {
				if parentNode == nil {
					tree.root = currentNode.Right
				} else {
					if currentNode.Value < parentNode.Value {
						parentNode.Left = currentNode.Right
					} else if currentNode.Value > parentNode.Value {
						parentNode.Right = currentNode.Right
					}
				}
			} else {
				// Find the Right child's Left most child
				leftMost := currentNode.Right.Left
				leftMostParent := currentNode.Right
				for leftMost.Left != nil {
					leftMostParent = leftMost
					leftMost = leftMost.Left
				}

				// Parent's Left subtree is now leftMost's Right subtree
				leftMostParent.Left = leftMost.Right
				leftMost.Left = currentNode.Left
				leftMost.Right = currentNode.Right

				if parentNode == nil {
					tree.root = leftMost
				} else {
					if currentNode.Value < parentNode.Value {
						parentNode.Left = leftMost
					} else if currentNode.Value > parentNode.Value {
						parentNode.Right = leftMost
					}
				}
			}
			return
		}
	}
}

/*
			9
		  /   \
		4		20
	   / \		/ \
	  1   6    15  170
*/

// Por não usarmos queue no DFS (Depth First Search),
// não temos a desvantagem de ter que armazenar todos os elementos em memória.
func (tree *BinarySearchTree) DFSInOrder() []int {
	return traverseInOrder(tree.root, []int{})
}

func traverseInOrder(node *Node, list []int) []int {
	if node.Left != nil {
		list = traverseInOrder(node.Left, list)
	}
	list = append(list, node.Value)
	if node.Right != nil {
		list = traverseInOrder(node.Right, list)
	}
	return list
}

func (tree *BinarySearchTree) DFSPreOrder() []int {
	return traversePreOrder(tree.root, []int{})
}

func traversePreOrder(node *Node, list []int) []int {
	list = append(list, node.Value)
	if node.Left != nil {
		list = traversePreOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePreOrder(node.Right, list)
	}
	return list
}

func (tree *BinarySearchTree) DFSPostOrder() []int {
	return traversePostOrder(tree.root, []int{})
}

func traversePostOrder(node *Node, list []int) []int {
	if node.Left != nil {
		list = traversePostOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePostOrder(node.Right, list)
	}
	list = append(list, node.Value)
	return list
}

func main() {
	myBST := NewBinarySearchTree()
	myBST.Insert(9)
	myBST.Insert(4)
	myBST.Insert(6)
	myBST.Insert(20)
	myBST.Insert(170)
	myBST.Insert(15)
	myBST.Insert(1)
	myBST.Print()

	fmt.Println(myBST.DFSInOrder())   // [1 4 6 9 15 20 170]
	fmt.Println(myBST.DFSPreOrder())  // [9 4 1 6 20 15 170]
	fmt.Println(myBST.DFSPostOrder()) // [1 6 4 15 170 20 9]

}

// Print exibe a árvore binária no formato JSON.
func (tree *BinarySearchTree) Print() {
	if tree.root == nil {
		fmt.Println("A árvore está vazia.")
		return
	}

	jsonData, err := json.MarshalIndent(tree.root, "", "  ") // Formata o JSON com identação
	if err != nil {
		fmt.Printf("Erro ao gerar JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData)) // Imprime o JSON formatado
}
