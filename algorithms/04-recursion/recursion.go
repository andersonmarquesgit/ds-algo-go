/*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)



Example 1:

Input: head = [1,2,3,4]

Output: [2,1,4,3]

Example 2:

Input: head = []

Output: []

Example 3:

Input: head = [1]

Output: [1]

Example 4:

Input: head = [1,2,3]

Output: [2,1,3]
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next

	first.Next = swapPairs(second.Next)
	second.Next = first

	return second
}

func main() {
	// Example 1
	// Input: head = [1,2,3,4]
	// Output: [2,1,4,3]
	// Explanation: The list is 1 -> 2 -> 3 -> 4
	// After swapping the nodes, the list becomes 2 -> 1 -> 4 -> 3
	// The function should return the head of the new list
	// 1 -> 2 -> 3 -> 4
	// 2 -> 1 -> 4 -> 3
	// head = 1 -> 2 -> 3 -> 4
	// head = 2 -> 1 -> 4 -> 3

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	newHead := swapPairs(head)
	fmt.Println(newHead) // 2

	head2 := &ListNode{Val: 1, Next: nil}
	newHead2 := swapPairs(head2)
	fmt.Println(newHead2) // 1
}
