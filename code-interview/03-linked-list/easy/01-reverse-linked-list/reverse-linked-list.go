package main

import linked_list "ds-algo-go/data-structures/util/linked-list"

/*
Time Complexity: O(n)
Space Complexity: O(1)
*/
func reverseLinkedList(head *linked_list.Node) *linked_list.Node {
	var prev, next *linked_list.Node
	currentNode := head

	for currentNode != nil {
		next = currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}

	return prev
}

func main() {
	linkedList := linked_list.NewLinkedList()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)

	linkedList.Print()
	linkedList.PrintByNode(reverseLinkedList(linkedList.Head))
}
