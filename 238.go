package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftArr := make([]int, n)
	rightArr := make([]int, n)
	leftArr[0] = 1
	rightArr[n-1] = 1
	for i := 1; i < n; i++ {
		leftArr[i] = nums[i-1] * leftArr[i-1]
		rightArr[n-i-1] = nums[n-i] * rightArr[n-i]
	}

	retArr := make([]int, n)
	retArr[0] = rightArr[0]
	retArr[n-1] = leftArr[n-1]
	for i := 1; i < n-1; i++ {
		retArr[i] = leftArr[i] * rightArr[i]
	}
	return retArr
}
