package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currNode := queue[0]

		for i := 0; i < levelSize; i++ {
			currNode = queue[0]
			queue = queue[1:]

			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}

			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}

		result = append(result, currNode.Val)
	}

	return result
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
	// Output: [3, 1, 4, 5, 8]

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

	fmt.Println(rightSideViewBFS(root)) // Expected:  [3, 1, 4, 5, 8]
}
