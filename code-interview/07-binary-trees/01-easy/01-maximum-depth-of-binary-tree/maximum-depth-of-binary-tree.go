package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
}

func calcMaximumDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := calcMaximumDepth(node.Left)
	right := calcMaximumDepth(node.Right)

	return max(left, right) + 1
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}

	return val2
}

func main() {
	// Test case 1
	//       *
	//      / \
	//   nil   nil
	// Output: 1

	tree1 := &TreeNode{
		Left:  nil,
		Right: nil,
	}
	fmt.Println(calcMaximumDepth(tree1)) // Expected: 1

	// Test case 2
	//            *
	//         /     \
	//        *        *
	//       / \     /   \
	//      *  nil  nil  nil
	//    /   \
	//  nil    *
	//       /   \
	//     nil    *
	//          /   \
	//        nil   nil
	// Output: 1

	tree2 := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: nil,
				Right: &TreeNode{
					Left: nil,
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(calcMaximumDepth(tree2)) // Expected: 5

	// Test case 2
	//      *
	//    /   \
	//  nil    *
	//       /   \
	//     nil    *
	//          /   \
	//       nil      *
	//              /   \
	//            nil    *
	//                 /   \
	//               nil    nil
	// Output: 1

	tree3 := &TreeNode{
		Left: nil,
		Right: &TreeNode{
			Left: nil,
			Right: &TreeNode{
				Left: nil,
				Right: &TreeNode{
					Left: nil,
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	fmt.Println(calcMaximumDepth(tree3)) // Expected: 5
}
