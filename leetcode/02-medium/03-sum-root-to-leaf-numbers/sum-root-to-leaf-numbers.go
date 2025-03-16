package main

import "fmt"

/*
129. Sum Root to Leaf Numbers
Medium
Topics
Companies
You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.



Example 1:
     1
    / \
   2   3

Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.

Example 2:
     4
    / \
   9   0
  / \
 5   1

Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.


Constraints:

The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 9
The depth of the tree will not exceed 10.
*/

/**
* Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	left, right := 0, 0
	if root.Left != nil {
		left = rootToLeaf(root.Left, root.Val)
	}
	if root.Right != nil {
		right = rootToLeaf(root.Right, root.Val)
	}

	return left + right
}

func rootToLeaf(node *TreeNode, rootVal int) int {
	if node.Left == nil && node.Right == nil {
		return node.Val + rootVal*10
	}

	rootValLevel := rootVal*10 + node.Val
	left, right := 0, 0

	if node.Left != nil {
		left = rootToLeaf(node.Left, rootValLevel)
	}

	if node.Right != nil {
		right = rootToLeaf(node.Right, rootValLevel)
	}

	return left + right
}

func main() {
	// Test case 1
	//     1
	//    / \
	//   2   3
	// Output: 25

	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(sumNumbers(tree1)) // Expected: 25

	// Test case 2
	//     4
	//    / \
	//   9   0
	//  / \
	// 5   1
	// Output: 1026

	tree2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(sumNumbers(tree2)) // Expected: 1026

	// Test case 3
	//     1
	//    / \
	//   0   0
	// Output: 20

	tree3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(sumNumbers(tree3)) // Expected: 20
}
