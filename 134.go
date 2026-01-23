package main

func canCompleteCircuit(gas []int, cost []int) int {
	res, total := 0, 0
	var ans int
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		res += gas[i] - cost[i]
		if res < 0 {
			ans = i + 1
			res = 0
		}
	}
	if total < 0 {
		return -1
	}
	return ans
}
