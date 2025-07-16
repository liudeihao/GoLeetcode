package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var ans int
	num2map := func(nums []int) map[int]int {
		ns := make(map[int]int)
		for _, n := range nums {
			ns[n]++
		}
		return ns
	}
	n1s, n2s, n3s, n4s := num2map(nums1), num2map(nums2), num2map(nums3), num2map(nums4)

	s1, s2 := make(map[int]int), make(map[int]int)
	for n1, c1 := range n1s {
		for n2, c2 := range n2s {
			s1[n1+n2] += c1 * c2
		}
	}
	for n3, c3 := range n3s {
		for n4, c4 := range n4s {
			s2[n3+n4] += c3 * c4
		}
	}

	for n1, c1 := range s1 {
		ans += c1 * s2[-n1]
	}

	return ans
}
