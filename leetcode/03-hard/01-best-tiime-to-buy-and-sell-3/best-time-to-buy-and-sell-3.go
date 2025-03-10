package main

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete at most two transactions.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).



Example 1:

Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time.
You must sell before buying again.
Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.


Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 105
*/
import (
	"fmt"
	"math"
)

// Função para calcular o máximo lucro com uma transação em um intervalo
func maxProfitSingleTransaction(prices []int) int {
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

func maxProfit(prices []int) int {
	n := len(prices)
	maxProfit := 0

	// Divide o array em dois pontos possíveis (antes e depois do dia j)
	for j := 0; j < n; j++ {
		// Calcula o lucro da primeira transação (antes de j)
		profit1 := maxProfitSingleTransaction(prices[:j])

		// Calcula o lucro da segunda transação (depois de j)
		profit2 := maxProfitSingleTransaction(prices[j:])

		// Atualiza o lucro máximo considerando essa divisão do array
		if profit1+profit2 > maxProfit {
			maxProfit = profit1 + profit2
		}
	}

	return maxProfit
}

func maxProfitOptimized(prices []int) int {
	firstBuy, secondBuy := math.MaxInt, math.MaxInt
	firstProfit, secondProfit := 0, 0

	for _, price := range prices {
		if price < firstBuy {
			firstBuy = price
		}

		if price-firstBuy > firstProfit {
			firstProfit = price - firstBuy
		}

		if price-firstProfit < secondBuy {
			secondBuy = price - firstProfit
		}

		if price-secondBuy > secondProfit {
			secondProfit = price - secondBuy
		}
	}

	return secondProfit
}

func main() {
	prices1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfitOptimized(prices1)) // 6
	fmt.Println(maxProfit(prices1))          // 6

	prices2 := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfitOptimized(prices2)) // 4
	fmt.Println(maxProfit(prices2))          // 4

	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfitOptimized(prices3)) // 0
	fmt.Println(maxProfit(prices3))          // 0

	prices4 := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	fmt.Println(maxProfitOptimized(prices4)) // 13
	fmt.Println(maxProfit(prices4))          // 13
}
