package main

func maxSubArray(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	ans := sum[1]
	min_sum := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, sum[i]-min_sum)
		min_sum = min(min_sum, sum[i])
	}
	return ans
}

func main() {
	maxSubArray([]int{5, 4, -1, 7, 8})
}
