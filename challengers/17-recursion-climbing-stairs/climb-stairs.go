/*
70. Climbing Stairs
Easy
Topics
Companies
Hint
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?



Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:

1 <= n <= 45
*/

package main

import "fmt"

func climbStairs(n int) int {
	/*
		For n less than 3, the function returns n directly because:
		If n is 1, there is only one way to climb the stairs (1 step).
		If n is 2, there are two ways to climb the stairs (1 step + 1 step or 2 steps).
	*/
	if n < 3 { // Base case
		return n
	}

	/*
		For values of n greater than or equal to 3, the function uses recursion
		to calculate the number of ways to climb
		the stairs by summing the results of the function called with n-1 and n-2:

		This recursive approach is based on the observation that the number of ways
		to reach the n-th step is the sum of the ways to reach the (n-1)-th step
		and the (n-2)-th step. This is because from the (n-1)-th step, you can take
		one step to reach the n-th step,
		and from the (n-2)-th step, you can take two steps to reach the n-th step.
	*/
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsInteractive(n int) int {
	if n == 1 {
		return n
	}

	climb := make([]int, n)
	climb[0], climb[1] = 1, 2
	for i := 2; i < n; i++ {
		climb[i] = climb[i-1] + climb[i-2]
	}
	return climb[n-1]
}

func main() {
	fmt.Println(climbStairs(2)) // 2
	fmt.Println(climbStairs(3)) // 3
	fmt.Println(climbStairs(4)) // 5
	fmt.Println(climbStairs(5)) // 8

	fmt.Println(climbStairsInteractive(2)) // 2
	fmt.Println(climbStairsInteractive(3)) // 3
	fmt.Println(climbStairsInteractive(4)) // 5
	fmt.Println(climbStairsInteractive(5)) // 8
}
