package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		cur := (r - l) * min(height[l], height[r])
		ans = max(ans, cur)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
