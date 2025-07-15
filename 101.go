package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isSame func(l, r *TreeNode) bool
	isSame = func(l, r *TreeNode) bool {
		ln, rn := l == nil, r == nil
		if ln && rn {
			return true
		}
		if (ln && !rn) || (!ln && rn) {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		return isSame(l.Left, r.Left) && isSame(l.Right, r.Right)
	}
	var flip func(*TreeNode)
	flip = func(node *TreeNode) {
		if node == nil {
			return
		}
		flip(node.Left)
		flip(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	flip(root.Right)

	return isSame(root.Left, root.Right)

}
