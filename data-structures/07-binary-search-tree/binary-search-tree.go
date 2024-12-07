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

func (tree *BinarySearchTree) Lookup(value int) *Node {
	return nil
}

func (tree *BinarySearchTree) Delete(value int) bool {
	return false
}

/*
			9
		  /   \
		4		20
	   / \		/ \
	  1   6    15  170
*/

func main() {
	myBST := NewBinarySearchTree()
	myBST.Insert(9)
	myBST.Insert(4)
	myBST.Insert(6)
	myBST.Insert(20)
	myBST.Insert(170)
	myBST.Insert(15)
	myBST.Insert(1)
	myBST.Insert(180)
	myBST.Insert(175)
	myBST.Print()
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

//// Print desenha a árvore binária com distinção entre raiz, nós à esquerda e à direita.
//func (tree *BinarySearchTree) Print() {
//	if tree.root == nil {
//		fmt.Println("A árvore está vazia.")
//		return
//	}
//
//	fmt.Printf("ROOT: %d\n", tree.root.Value)   // Identifica a raiz
//	printTree(tree.root.Left, "      ", true)   // Nós à esquerda da raiz
//	printTree(tree.root.Right, "      ", false) // Nós à direita da raiz
//}
//
//// Função auxiliar para imprimir a árvore com distinção entre nós à esquerda e à direita
//func printTree(node *Node, prefix string, isLeft bool) {
//	if node != nil {
//		if isLeft {
//			fmt.Printf("%s└──(L) %d\n", prefix, node.Value)
//		} else {
//			fmt.Printf("%s└──(R) %d\n", prefix, node.Value)
//		}
//
//		newPrefix := prefix + "    "
//		printTree(node.Left, newPrefix+"    ", true)
//		printTree(node.Right, newPrefix+"    ", false)
//	}
//}
