package main

func canCompleteCircuit(gas []int, cost []int) int {
	for start, n := 0, len(gas); start < n; {
		gasSum, costSum := 0, 0
		step := 0
		for step < n {
			if step == n {
				return start
			}
			idx := (start + step) % n
			gasSum += gas[idx]
			costSum += cost[idx]

			if costSum > gasSum {
				break
			}
			step++
		}
		if step == n {
			return start
		} else {
			start = start + step + 1
		}
	}
	return -1
}
