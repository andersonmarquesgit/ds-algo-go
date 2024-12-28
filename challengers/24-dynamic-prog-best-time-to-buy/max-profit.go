/*
121. Best Time to Buy and Sell Stock
Easy
Topics
Companies
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.



Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.


Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
*/

package main

import "fmt"

// Not using dynamic programming
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

// Using dynamic programming
func maxProfit2(prices []int) int {
	cache := make(map[int]int)

	return maxProfitRecursive(prices, 0, cache)
}

func maxProfitRecursive(prices []int, i int, cache map[int]int) int {
	if i >= len(prices) {
		return 0
	}

	if val, ok := cache[i]; ok {
		return val
	}

	buy := prices[i]
	maxProfit := 0

	for j := i + 1; j < len(prices); j++ {
		sell := prices[j]
		profit := sell - buy

		if profit > 0 {
			profit += maxProfitRecursive(prices, j+1, cache)
		}

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	cache[i] = maxProfit
	return cache[i]
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))  // 5
	fmt.Println(maxProfit2(prices)) // 5

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))  // 0
	fmt.Println(maxProfit2(prices)) // 0
}
