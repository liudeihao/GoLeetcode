package main

func maxProfit(prices []int) int {
	pre := prices[0]
	ret := 0
	for _, cur := range prices {
		if cur > pre {
			ret += cur - pre
		}
		pre = cur
	}
	return ret
}
