package main

import (
	"fmt"
	"math"
)

/*
121. Best Time to Buy and Sell Stock
Solved
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

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices1))

	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices2))

	prices3 := []int{2, 1, 2, 0, 1}
	fmt.Println(maxProfit(prices3))

	prices4 := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(prices4))
}

// [7,1,5,3,6,4] ----> 5
// [7,6,4,3,1]   ----> 0
// [2,1,4]       ----> 3
// [2,4,1]       ----> 2
// [1,4,2]       ----> 3
// [2,1,2,0,1]   ----> 1
// [3,2,6,5,0,3] ----> 4
