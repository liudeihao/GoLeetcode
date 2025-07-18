package main

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var traverse func(r *TreeNode)
	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}
		traverse(r.Left)
		ans = append(ans, r.Val)
		traverse(r.Right)
	}
	traverse(root)
	return ans
}
