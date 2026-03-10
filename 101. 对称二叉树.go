package main

func isSymmetric(root *TreeNode) bool {
	var isMirror func(l, r *TreeNode) bool
	isMirror = func(l, r *TreeNode) bool {
		// 比较两棵树是否镜像
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		return l.Val == r.Val && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
	}
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}
