package main

// 更简单的方法是让左子树整个放到右子树最左下的位置！
// 应当想到那个方法
func deleteRoot(r *TreeNode) *TreeNode {
	if r == nil {
		return nil
	}
	if r.Left == nil && r.Right == nil {
		return nil
	}
	if r.Left != nil {
		p := r.Left
		if p.Right == nil {
			p.Right = r.Right
			return p
		}
		prev := r
		for p.Right != nil {
			prev = p
			p = p.Right
		}
		prev.Right = p.Left
		p.Left = r.Left
		p.Right = r.Right
		return p
	} else {
		p := r.Right
		if p.Left == nil {
			p.Left = r.Left
			return p
		}
		prev := r
		for p.Left != nil {
			prev = p
			p = p.Left
		}
		prev.Left = p.Right
		p.Left = r.Left
		p.Right = r.Right

		return p
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		return deleteRoot(root)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
