package main

func countValidSelections(nums []int) int {
	// 打砖块，可以用前缀和
	n := len(nums)
	sum := 0
	// 右前缀和就是sum-左前缀和
	for _, x := range nums {
		sum += x
	}
	leftSum, rightSum := 0, sum
	ans := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			if leftSum == rightSum {
				ans += 2
			} else if leftSum == rightSum+1 || leftSum+1 == rightSum {
				ans++
			}
		} else {
			leftSum += nums[i]
			rightSum -= nums[i]
		}
	}
	return ans
}

/*
模拟。
func countValidSelections(nums []int) int {

	n := len(nums)

	tryit := func(start int, rightward bool) int {
		ns := make([]int, n)
		copy(ns, nums)
		l, r := start, start
		for l >= 0 && r < n {
			for l >= 0 && ns[l] == 0 {
				l--
			}
			for r < n && ns[r] == 0 {
				r++
			}
			if rightward && r < n {
				ns[r]--
				rightward = false
			} else if !rightward && l >= 0 {
				ns[l]--
				rightward = true
			}
		}
		for l >= 0 && ns[l] == 0 {
			l--
		}
		for r < n && ns[r] == 0 {
			r++
		}
		if l < 0 && r >= n {
			return 1
		}
		return 0
	}

	ans := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			continue
		}
		ans += tryit(i, true)
		ans += tryit(i, false)
	}
	return ans
}
*/
