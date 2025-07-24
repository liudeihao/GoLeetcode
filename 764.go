package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	minCost := make([]int, n+1)
	minCost[n] = 0
	minCost[n-1] = cost[n-1]
	minCost[n-2] = cost[n-2]
	for i := n - 3; i >= 0; i-- {
		minCost[i] = cost[i] + min(minCost[i+1], minCost[i+2])
	}
	return min(minCost[0], minCost[1])
}
