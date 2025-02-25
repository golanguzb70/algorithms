package main

/*Problem:
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

/*Solution:
Algorithm: Sliding window or Two pointer
p1=0, p2=1
Extend the window to the right if current element is greater than or equal to last profit
Shrink the window from the left if current element is less than last bought stock

Time complexity: O(n)
Space complexity: O(1)
*/

func MaxProfit(prices []int) int {
	response := 0
	boughtDay := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[boughtDay] {
			boughtDay = i
			continue
		}

		if prices[i]-prices[boughtDay] > response {
			response = prices[i] - prices[boughtDay]
		}

	}

	return response
}
