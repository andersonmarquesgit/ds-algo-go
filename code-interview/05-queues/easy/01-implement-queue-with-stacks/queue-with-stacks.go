package main

import "fmt"

/*
232. Implement MyQueue using Stacks
Easy
Topics
Companies
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty
operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque
(double-ended queue) as long as you use only a stack's standard operations.

Example 1:

Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false

Constraints:

1 <= x <= 9
At most 100 calls will be made to push, pop, peek, and empty.
All the calls to pop and peek are valid.

Follow-up: Can you implement the queue such that each operation is amortized O(1) time complexity? In other words, performing n operations will take overall O(n) time even if one of those operations may take longer.
*/
type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  nil,
		stackOut: nil,
	}
}

// queue
func (q *MyQueue) push(value int) { // Tp e Sp: O(1)
	q.stackIn = append(q.stackIn, value)
}

// dequeue
func (q *MyQueue) pop() int { // Tp e Sp: O(n)
	if len(q.stackOut) == 0 {
		for len(q.stackIn) > 0 {
			q.stackOut = append(q.stackOut, q.stackIn[len(q.stackIn)-1])
			q.stackIn = q.stackIn[:len(q.stackIn)-1]
		}
	}

	valuePop := q.stackOut[len(q.stackOut)-1]
	q.stackOut = q.stackOut[:len(q.stackOut)-1]
	return valuePop
}

func (q *MyQueue) empty() bool { // Tp e Sp: O(1)
	return len(q.stackIn) == 0 && len(q.stackOut) == 0
}

func (q *MyQueue) peek() int { // Tp e Sp: O(n)
	if len(q.stackOut) == 0 {
		for len(q.stackIn) > 0 {
			q.stackOut = append(q.stackOut, q.stackIn[len(q.stackIn)-1])
			q.stackIn = q.stackIn[:len(q.stackIn)-1]
		}
	}

	return q.stackOut[len(q.stackOut)-1]
}

func main() {
	myQueue := Constructor()
	myQueue.push(1)
	myQueue.push(2)
	fmt.Println(myQueue.stackIn)
	fmt.Println(myQueue.stackOut)

	fmt.Println("------")

	myQueue.pop()
	fmt.Println(myQueue.stackIn)
	fmt.Println(myQueue.stackOut)

	fmt.Println("------")

	fmt.Println(myQueue.peek())
	fmt.Println(myQueue.stackIn)
	fmt.Println(myQueue.stackOut)

	fmt.Println("------")

	myQueue.push(3)
	myQueue.push(4)
	fmt.Println(myQueue.stackIn)
	fmt.Println(myQueue.stackOut)

	fmt.Println("------")

	myQueue.pop()
	fmt.Println(myQueue.stackIn)
	fmt.Println(myQueue.stackOut)
}
