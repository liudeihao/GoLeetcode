package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p, q := m-1, n-1
	i := m + n - 1
	for ; p >= 0 && q >= 0; i-- {
		if nums1[p] >= nums2[q] {
			nums1[i] = nums1[p]
			p--
		} else {
			nums1[i] = nums2[q]
			q--
		}
	}
	for ; i >= 0 && q >= 0; i-- {
		nums1[i] = nums2[q]
		q--
	}
}
