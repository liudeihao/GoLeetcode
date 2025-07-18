package main

func isSameTree(l, r *TreeNode) bool {
	ln, rn := l == nil, r == nil
	if ln && rn {
		return true
	}
	if ln || rn { // 确实
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return isSameTree(l.Left, r.Left) && isSameTree(l.Right, r.Right)
}
