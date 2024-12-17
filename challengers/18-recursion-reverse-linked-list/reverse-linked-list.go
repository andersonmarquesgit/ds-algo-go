/*Given the head of a singly linked list, reverse the list, and return the reversed list.



Example 1:


Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
Example 2:


Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []


Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000


Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // Base case
		return head
	}

	// For example, if the list is 1 -> 2 -> 3 -> 4 -> 5
	// The recursive call will reverse the list from 2 to 5
	// The head
	// 1 -> 2 -> 3 -> 4 -> 5
	// The reversed list
	// 5 -> 4 -> 3 -> 2 -> 1
	p := reverseListRecursion(head.Next) // Recursive call
	head.Next.Next = head
	head.Next = nil

	return p

}

func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func main() {
	// Example 1
	// Input: head = [1,2,3,4,5]
	// Output: [5,4,3,2,1]
	// Explanation: The list is 1 -> 2 -> 3 -> 4 -> 5
	// After reversing the list, the list becomes 5 -> 4 -> 3 -> 2 -> 1
	// The function should return the head of the new list
	// 1 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reversed := reverseListRecursion(head)
	for reversed != nil {
		fmt.Print(reversed.Val)
		if reversed.Next != nil {
			fmt.Print(" -> ")
		}
		reversed = reversed.Next
	}
}
