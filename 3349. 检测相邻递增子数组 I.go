package main

func hasIncreasingSubarrays(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	for i := 0; i+2*k <= len(nums); {
		ok := true
		for j := 0; j < k-1; j++ {
			if nums[i+j] >= nums[i+j+1] || nums[i+k+j] >= nums[i+k+j+1] {
				i = i + j + 1
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
