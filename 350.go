package main

func intersect(nums1 []int, nums2 []int) []int {
	s1, s2 := make(map[int]int), make(map[int]int)
	for _, n := range nums1 {
		s1[n]++
	}
	for _, n := range nums2 {
		s2[n]++

	}
	var ans []int
	for k, v := range s2 {
		cnt := min(v, s1[k])
		for cnt > 0 {
			cnt--
			ans = append(ans, k)
		}
	}
	return ans
}
