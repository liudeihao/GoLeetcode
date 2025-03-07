package main

func maxProfit(prices []int) int {
	profit := 0
	lowest := 100_000
	for _, price := range prices {
		lowest = min(lowest, price)
		profit = max(profit, price-lowest)
	}
	return profit
}
