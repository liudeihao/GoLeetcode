package main

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var traverse func(r *TreeNode)
	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}
		traverse(r.Left)
		traverse(r.Right)
		ans = append(ans, r.Val)
	}
	traverse(root)
	return ans
}
