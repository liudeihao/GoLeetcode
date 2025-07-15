package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for i, num := range nums {
		prev, ok := mp[num]
		if ok && i-prev <= k {
			return true
		}
		mp[num] = i
	}
	return false
}
