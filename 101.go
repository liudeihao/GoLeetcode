package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isSym func(l, r *TreeNode) bool
	isSym = func(l, r *TreeNode) bool {
		ln, rn := l == nil, r == nil
		if ln && rn {
			return true
		}
		if ln || rn {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		// 嗯，这样是正确的。
		return isSym(l.Right, r.Left) && isSym(l.Left, r.Right)
	}

	return isSym(root.Left, root.Right)

}
