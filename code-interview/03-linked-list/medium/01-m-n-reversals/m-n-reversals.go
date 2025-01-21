package main

import linked_list "ds-algo-go/data-structures/util/linked-list"

/*
Time Complexity: O(n)
Space Complexity: O(1)
*/
func reversals(head *linked_list.Node, m, n int) *linked_list.Node {
	var newList, next *linked_list.Node
	currentNode := head
	start := head
	position := 1

	for position < m {
		start = currentNode
		currentNode = currentNode.Next
		position++
	}

	tail := currentNode

	for position >= m && position <= n {
		next = currentNode.Next
		currentNode.Next = newList
		newList = currentNode
		currentNode = next
		position++
	}

	start.Next = newList
	tail.Next = currentNode

	if m > 1 {
		return head
	} else {
		return newList
	}
}

func main() {
	linkedList := linked_list.NewLinkedList()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	linkedList.Append(6)
	linkedList.Append(7)

	linkedList.Print()
	linkedList.PrintByNode(reversals(linkedList.Head, 3, 5))

	linkedList2 := linked_list.NewLinkedList()
	linkedList2.Append(1)

	linkedList2.Print()
	linkedList2.PrintByNode(reversals(linkedList2.Head, 1, 1))
}
