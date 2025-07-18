package main

func sumOfLeftLeaves(r *TreeNode) int {
	if r == nil {
		return 0
	}
	if r.Left != nil && r.Left.Left == nil && r.Left.Right == nil {
		return r.Left.Val + sumOfLeftLeaves(r.Right)
	}
	return sumOfLeftLeaves(r.Left) + sumOfLeftLeaves(r.Right)
}
