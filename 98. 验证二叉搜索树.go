package main

func isValidBST(root *TreeNode) bool {
	// 练习一下迭代遍历
	var s []*TreeNode
	var prev *TreeNode
	cur := root
	for cur != nil || len(s) > 0 {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}

		cur = s[len(s)-1]
		s = s[:len(s)-1]

		// 中序遍历
		if prev != nil && cur.Val <= prev.Val {
			return false
		}
		prev = cur

		cur = cur.Right
	}
	return true
}

/*
// 递归
func isValidBST(root *TreeNode) bool {
	var res []int
	var trav func(node *TreeNode)
	trav = func(node *TreeNode) {
		if node == nil {
			return
		}
		trav(node.Left)
		res = append(res, node.Val)
		trav(node.Right)
	}
	trav(root)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

*/
