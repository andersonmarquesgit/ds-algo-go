package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderOfBinaryTree(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var order [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currArr := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currArr = append(currArr, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		order = append(order, currArr)
	}

	return order
}

func main() {
	// Test case 1
	//           3
	//         /   \
	//        6     1
	//      /   \    \
	//     9     2    4
	//    /  \
	//  nil   5
	//       /  \
	//      8    nil
	//
	//
	// Output: [[3] [6 1] [9 2 4] [5] [8]]

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:  9,
				Left: nil,
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(levelOrderOfBinaryTree(root)) // Expected: 1
}
