package main

// 3652. Best Time to Buy and Sell Stock using Strategy
// You are given two integer arrays prices and strategy, where:
//    prices[i] is the price of a given stock on the ith day.
//    strategy[i] represents a trading action on the ith day, where:
//        -1 indicates buying one unit of the stock.
//        0 indicates holding the stock.
//        1 indicates selling one unit of the stock.
//
// You are also given an even integer k, and may perform at most one modification to strategy. A modification consists of:
//    Selecting exactly k consecutive elements in strategy.
//    Set the first k / 2 elements to 0 (hold).
//    Set the last k / 2 elements to 1 (sell).
//
// The profit is defined as the sum of strategy[i] * prices[i] across all days.
// Return the maximum possible profit you can achieve.
// Note: There are no constraints on budget or stock ownership, so all buy and sell operations are feasible regardless
// of past actions.

func maxProfit(prices []int, strategy []int, k int) int64 {
	// not modified
	profit := sum(prices, strategy)
	maxProf := profit
	halfK := k / 2
	//first modified
	for i := 0; i < k; i++ {
		if i < halfK {
			// minus operation
			profit -= strategy[i] * prices[i]
		} else {
			// replace operation
			profit += (1 - strategy[i]) * prices[i]
		}
	}
	maxProf = max(maxProf, profit)

	for i := 1; i <= len(prices)-k; i++ {
		// revert previous operation
		profit += prices[i-1] * strategy[i-1]
		// remove modification breakpoint
		profit -= prices[i+halfK-1]
		// replace last operation by mod
		profit += (1 - strategy[i+k-1]) * prices[i+k-1]
		maxProf = max(maxProf, profit)
	}

	return int64(maxProf)
}

func sum(prices []int, strategy []int) (s int) {
	for i, p := range prices {
		if strategy[i] == 0 {
			continue
		} else if strategy[i] == 1 {
			s += p
		} else {
			s -= p
		}
	}
	return s
}

/**
k=2 -1     0    1    2    3    4
4 = orig + 01 + 12 + 23 = 4 = len = 2len / k = 1 + (len - k + 1)
5 = orig + 01 + 12 + 23 + 34 = 5 = len
6 = orig + 01 + 12 + 23 + 34 + 45 = 6 = len
*/

/**
k=4 -1     0      1      2
4 = orig + 0123 = len / 2 = 2 len / k
5 = orig + 0123 + 1234 = 3 = 1 + (len - k + 1)
6 = orig + 0123 + 1234 + 2345 = 4 = 1 + (len - k + 1)
*/
