package main

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var trav func(*TreeNode)
	trav = func(r *TreeNode) {
		if r == nil {
			return
		}
		trav(r.Right)
		sum += r.Val
		r.Val = sum
		trav(r.Left)
	}
	trav(root)
	return root
}
