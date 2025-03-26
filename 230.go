package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	cur := 0
	ans := -1
	var midorder func(*TreeNode)
	midorder = func(r *TreeNode) {
		if ans >= 0 || r == nil {
			return
		}
		midorder(r.Left)
		cur++
		if cur == k {
			ans = r.Val
		}
		midorder(r.Right)
	}
	midorder(root)
	return ans
}
