package main

func intersection(nums1 []int, nums2 []int) []int {
	s1, s2 := make(map[int]bool), make(map[int]bool)
	for _, n := range nums1 {
		s1[n] = true
	}
	for _, n := range nums2 {
		s2[n] = true

	}
	var ans []int
	for k := range s2 {
		if s1[k] {
			ans = append(ans, k)
		}
	}
	return ans
}
