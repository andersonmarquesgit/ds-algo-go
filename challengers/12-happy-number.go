package main

/**
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:

Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
Example 2:

Input: n = 2
Output: false
*/

func isHappy(n int) bool {
	slow, fast := n, sumOfSquares(n)
	for slow != fast {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
	}
	return slow == 1
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10  // digit
		sum += d * d // square of digit
		n /= 10      // remove digit
	}
	return sum
}

func main() {
	println(isHappy(19)) // true
	println(isHappy(2))  // false
}
