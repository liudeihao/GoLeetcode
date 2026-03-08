package main

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1
			nums[j], nums[i] = nums[i], nums[j]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

/*
// 带注释的。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 这道题的本质就是“让数字 x 回到索引 x-1 的位置上”

	for i := 0; i < n; i++ {
		// 只要当前位置上的数字 x 是有效的 (1 <= x <= n)
		// 并且 x 没在它该在的位置上 (nums[x-1] 还不等于 x)
		// 这个 nums[nums[i]-1] != nums[i] 是给 nums[i] - 1 != i 加了一层。
		// “目标位置上是不是已经躺着正确的数字了？”, 它可以解决重复元素的情况.
		// 如果数组是 [1, 1]，如果没有这个条件，数字 1 会不断地和另一个数字 1 交换，直接死循环。
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 把当前的数字 nums[i] 扔回它的老家索引 nums[i]-1
			// 同时把老家那个位置原来的数字换回来，继续在 for 里检查它
			j := nums[i] - 1
			nums[j], nums[i] = nums[i], nums[j]
		}
		// 内部 for 结束时，要么 nums[i] 是正确的数 (i+1)
		// 要么 nums[i] 是个无用的垃圾值（负数/超大数/重复数），不用管它，继续看下一个 i
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
*/
