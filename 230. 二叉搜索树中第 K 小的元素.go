package main

func kthSmallest(root *TreeNode, k int) int {
	var s []*TreeNode

	cur := root
	for cur != nil || len(s) > 0 {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}

		cur = s[len(s)-1]
		s = s[:len(s)-1]

		// 中序遍历
		k--
		if k == 0 {
			return cur.Val
		}

		cur = cur.Right
	}
	return -1
}

/*
// 递归

func kthSmallest(root *TreeNode, k int) int {
	cur := 0
	ans := -1
	var midorder func(*TreeNode)
	midorder = func(r *TreeNode) {
		if ans >= 0 || r == nil {
			return
		}
		midorder(r.Left)
		cur++
		if cur == k {
			ans = r.Val
		}
		midorder(r.Right)
	}
	midorder(root)
	return ans
}

*/
