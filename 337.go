package main

func rob(root *TreeNode) int {
	var _rob func(root *TreeNode) (int, int)
	_rob = func(root *TreeNode) (steal, skip int) {
		// 返回1表示偷了，0表示没偷
		if root == nil {
			return 0, 0
		}
		lSteal, lSkip := _rob(root.Left)
		rSteal, rSkip := _rob(root.Right)
		steal = lSkip + rSkip + root.Val
		skip = max(lSteal, lSkip) + max(rSteal, rSkip)
		return steal, skip
	}
	steal, skip := _rob(root)
	return max(steal, skip)
}
