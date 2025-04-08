package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	l, r := 0, len(nums)
	m := (l + r) / 2
	return &TreeNode{
		nums[m],
		sortedArrayToBST(nums[l:m]),
		sortedArrayToBST(nums[m+1 : r]),
	}
}
