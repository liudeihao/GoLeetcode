package main

func partitionLabels(s string) []int {
	// 就是求不重叠区间。
	// 但具体到这个题，不用按模板的写法
	rbound := [26]int{}
	for i, c := range s {
		rbound[c-'a'] = i
	}

	var ans []int
	left, right := 0, 0
	for i, c := range s {
		right = max(right, rbound[c-'a']) // 字符出现的最远边界
		if right == i {
			ans = append(ans, right-left+1)
			left = i + 1
		}
	}
	return ans

}
