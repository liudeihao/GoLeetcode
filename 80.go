package main

func removeDuplicates(nums []int) int {
	// 要删除的nums[i]会满足nums[i]==nums[i-2]
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow := 2
	for fast := 2; fast < n; fast++ {
		// 这里是slow-2
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
