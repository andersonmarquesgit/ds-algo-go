/*
ou are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

Example 1:

Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]
Example 2:

Input: root = [4,2,7,1,3], val = 5
Output: []

Constraints:

The number of nodes in the tree is in the range [1, 5000].
1 <= Node.val <= 107
root is a binary search tree.
1 <= val <= 107
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val) // Recursive call on the left subtree
	}

	return searchBST(root.Right, val) // Recursive call on the right subtree
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// Example 1
	// Input: root = [4,2,7,1,3], val = 2
	// Output: [2,1,3]
	// Explanation: The node with the value 2 is found in the tree.
	// The subtree rooted with the node with the value 2 is [2,1,3]
	// The function should return the root of the subtree
	//    4
	//   / \
	//  2   7
	// / \
	// 1 3

	//  2
	// / \
	// 1 3
	// Example 2
	// Input: root = [4,2,7,1,3], val = 5
	// Output: []
	// Explanation: There is no node with the value 5 in the tree.
	// The function should return null
	//    4
	//   / \
	//  2   7
	// / \
	// 1 3
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	subtree := searchBST(tree, 2)
	printTree(subtree)
}
