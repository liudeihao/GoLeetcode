package main

func binarySearch(n int, cond func(int) bool) (index int) {
	l, r := 0, n // 右开区间定义是这样的
	for l < r {
		m := (l + r) / 2
		if cond(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	// 在两边且无重叠，直接返回
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	if newInterval[0] > intervals[n-1][1] {
		return append(intervals, newInterval)
	}
	le := binarySearch(n, func(mid int) bool {
		return intervals[mid][0] > newInterval[0]
	}) - 1
	ge := binarySearch(n, func(mid int) bool {
		return intervals[mid][0] > newInterval[1]
	}) - 1

	if le == -1 {
		newInterval[1] = max(intervals[ge][1], newInterval[1])
		return append([][]int{newInterval}, intervals[ge+1:]...)
	}

	if intervals[le][1] >= newInterval[0] {
		intervals[le][1] = max(intervals[ge][1], newInterval[1])
		return append(intervals[:le+1], intervals[ge+1:]...)
	} else {
		return append(intervals[:le+1], insert(intervals[le+1:], newInterval)...)
	}
}
