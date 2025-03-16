package main

import (
	"fmt"
	"math"
)

/*
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting
them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.



Example 1:


Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
Example 2:


Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42
*/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}

	// Obtém a soma máxima começando do nó atual
	maxFromRoot := maxPathSumFromNode(root)

	// Obtém a soma máxima considerando apenas os filhos
	leftMax := maxPathSum(root.Left)
	rightMax := maxPathSum(root.Right)

	// O maior valor considerando todas as opções
	return max(maxFromRoot, leftMax, rightMax)
}

func maxPathSumFromNode(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// Soma máxima incluindo este nó e os filhos
	leftSum := maxPathSumFromNode(node.Left)
	rightSum := maxPathSumFromNode(node.Right)

	return max(node.Val, node.Val+leftSum, node.Val+rightSum, node.Val+leftSum+rightSum)
}

func maxPathSumOptimized(root *TreeNode) int {
	maxSum := math.MinInt

	// Função recursiva para calcular a soma máxima de um caminho que desce a partir de um nó
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Calcula o máximo caminho da subárvore esquerda e direita
		leftMax := max2(dfs(node.Left), 0)   // Se for negativo, descartamos (consideramos 0)
		rightMax := max2(dfs(node.Right), 0) // Se for negativo, descartamos (consideramos 0)

		// Atualiza a soma máxima global considerando o caminho que passa pelo nó atual
		currentPathSum := node.Val + leftMax + rightMax
		maxSum = max2(maxSum, currentPathSum)

		// Retorna o melhor caminho **descendo** para continuar a recursão
		return node.Val + max2(leftMax, rightMax)
	}

	dfs(root)
	return maxSum
}

// Função auxiliar para obter o máximo entre dois valores
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Função auxiliar para calcular o máximo entre vários valores
func max(values ...int) int {
	maxValue := math.MinInt
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// Função para testar o código
func main() {
	// Teste 1: [1,2,3]
	root1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(maxPathSum(root1))          // 6 - O(n^2)
	fmt.Println(maxPathSumOptimized(root1)) // 6 - O(n)

	// Teste 2: [-10,9,20,null,null,15,7]
	root2 := &TreeNode{
		Val:   -10,
		Left:  &TreeNode{9, nil, nil},
		Right: &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	}
	fmt.Println(maxPathSum(root2))          // 42
	fmt.Println(maxPathSumOptimized(root2)) // 42
}
