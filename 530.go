package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	ans := 1_000_000
	prev := -1_000_000
	var midOrder func(r *TreeNode)
	midOrder = func(r *TreeNode) {
		if r == nil {
			return
		}
		midOrder(r.Left)
		if r.Val-prev < ans {
			ans = r.Val - prev
		}
		prev = r.Val
		midOrder(r.Right)
	}
	midOrder(root)
	return ans
}
