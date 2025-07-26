package main

func maxProfit(k int, prices []int) int {
	n := len(prices)
	buy := make([]int, k)
	sell := make([]int, k)
	for i := range buy {
		buy[i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])
		for j := 1; j < k; j++ {
			// 用的都是当天的，但没问题！
			// sell[j-1]取sell[j-1]时，buy[j-1]不变
			// sell[j-1]取buy[j-1]+prices[i]时，说明这天要卖，sell[j-1]-prices[i]==buy[j-1]不变
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			// buy[j-1]取buy[j-1]时，保持，所以sell[j]不变
			// buy[j-1]取sell[j-1]-prices[i]时，说明这天要买，buy[j-1]+prices[i]==sell[j-1]不变
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	// 最后一次交易一定最大（它包括了不足k次交易的情况）
	return sell[k-1]
}
