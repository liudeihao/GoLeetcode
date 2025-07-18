package main

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var traverse func(r *TreeNode)
	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}
		ans = append(ans, r.Val)
		traverse(r.Left)
		traverse(r.Right)
	}
	traverse(root)
	return ans
}
