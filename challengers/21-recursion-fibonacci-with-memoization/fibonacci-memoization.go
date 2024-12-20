/*
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).

Example 1:

Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
Example 2:

Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
Example 3:

Input: n = 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

Constraints:

0 <= n <= 30
*/

package main

import "fmt"

/*
Use memoization for recursion:
Memoization is an optimization technique used primarily to speed up computer programs by storing the results of
expensive function calls and returning the cached result when the same inputs occur again. (Source: wikipedia)

public class Main {

	  HashMap<Integer, Integer> cache = new HashMap<Integer, Integer>();

	  private int fib(int N) {
	    if (cache.containsKey(N)) {
	      return cache.get(N);
	    }
	    int result;
	    if (N < 2) {
	      result = N;
	    } else {
	      result = fib(N-1) + fib(N-2);
	    }
	    // keep the result in cache.
	    cache.put(N, result);
	    return result;
	  }
	}

Given a number N return the index value of the Fibonacci sequence, where the sequence is:
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
*/
func fib(n int) int {
	cache := make(map[int]int)

	var memoFib func(int) int
	memoFib = func(n int) int {
		if val, ok := cache[n]; ok { // Verifica se o valor já está em cache
			return val
		}

		if n < 2 { // Casos base
			cache[n] = n
			return n
		}

		// Calcula e armazena no cache
		cache[n] = memoFib(n-1) + memoFib(n-2)
		return cache[n]
	}

	return memoFib(n)
}

func main() {
	fmt.Println(fib(4))
	fmt.Println(fib(9))
}
