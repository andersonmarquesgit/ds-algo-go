package main

/*
Input (stdin)
7
3 5 2 1 4 6 7
Your Output (stdout)
3
Expected Output
3

Input (stdin)
5
3 1 7 5 4
Your Output (stdout)
2

Input (stdin)
1
15
Your Output (stdout)
0
Expected Output
0

*/
import (
	"encoding/json"
	"fmt"
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

func (tree *BinarySearchTree) Height() int {
	return heightRecursion(tree.root)
}

// heightRecursion calcula a altura da árvore binária de forma recursiva. O retorno é a altura da árvore.
func heightRecursion(node *TreeNode) int { // O(n)
	if node == nil {
		return -1
	}

	leftHeight := heightRecursion(node.Left)   // Calcula a altura do lado esquerdo da árvore recursivamente
	rightHeight := heightRecursion(node.Right) // Calcula a altura do lado direito da árvore recursivamente

	if leftHeight > rightHeight { // Verifica qual lado é maior e retorna a altura
		return leftHeight + 1 // Adiciona 1 para contar o nó atual
	}

	return rightHeight + 1
}

func main() {
	var n int
	fmt.Scan(&n)
	myBinarySearchTree := NewBinarySearchTree()

	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		myBinarySearchTree.Insert(value)
	}

	myBinarySearchTree.Print()
	fmt.Println("Height:", myBinarySearchTree.Height())
}
