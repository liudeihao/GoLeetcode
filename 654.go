package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal, maxIdx := -1, -1
	for i, v := range nums {
		if v > maxVal {
			maxVal, maxIdx = v, i
		}
	}
	return &TreeNode{
		Val:   maxVal,
		Left:  constructMaximumBinaryTree(nums[:maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
	}
}
