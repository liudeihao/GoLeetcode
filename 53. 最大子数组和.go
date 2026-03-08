package main

func maxSubArray(nums []int) int {
	// 所谓Kadane算法，就是做最大子数组和的。
	ans := nums[0]
	sum := 0

	for _, num := range nums {
		// 如果之前的累加和大于 0，说明对当前元素有增益，继续累加
		if sum > 0 {
			sum += num
		} else {
			// 如果之前的累加和小于等于 0，说明是拖累，直接从当前元素重新开始计算
			sum = num
		}
		ans = max(ans, sum)
	}
	return ans
}

/*
// 因为是连续子数组的和，所以用前缀和
func maxSubArray(nums []int) int {
	// 后续前缀和减去之前的最小的前缀和，就好了。
	minSum := 0
	sum := 0
	ans := math.MinInt
	for _, v := range nums {
		sum += v
		ans = max(ans, sum-minSum)
		minSum = min(minSum, sum)
	}
	return ans
}
*/
