package main

func trap(height []int) int {
	n := len(height)

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	lmax := 0
	for i := 0; i < n; i++ {
		leftMax[i] = lmax
		lmax = max(lmax, height[i])
	}
	rmax := 0
	for i := n - 1; i >= 0; i-- {
		rightMax[i] = rmax
		rmax = max(rmax, height[i])
	}
	result := 0
	for i := range height {
		result += max(min(leftMax[i], rightMax[i])-height[i], 0)
	}
	return result
}
