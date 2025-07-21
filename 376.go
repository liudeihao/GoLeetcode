package main

func wiggleMaxLength(nums []int) int {
	ans := 1
	isUp := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] && isUp != 1 {
			ans++
			isUp = 1
		}
		if nums[i] < nums[i-1] && isUp != -1 {
			ans++
			isUp = -1
		}
	}
	return ans
}
