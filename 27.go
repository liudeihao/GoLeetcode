package main

func removeElement(nums []int, val int) int {
	// 快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组
	// 慢指针：指向更新 新数组下标的位置
	n := len(nums)
	slow := 0
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
