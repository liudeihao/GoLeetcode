package main

func isBalanced(root *TreeNode) bool {
	ans := true
	var isBal func(*TreeNode) (depth int, bal bool)
	isBal = func(r *TreeNode) (int, bool) {
		if ans == false {
			return -1, false
		}
		if r == nil {
			return 0, true
		}
		ld, lb := isBal(r.Left)
		rd, rb := isBal(r.Right)
		if ld > rd+1 || rd > ld+1 || !lb || !rb {
			ans = false
			return -1, false
		}
		return max(ld, rd) + 1, true
	}
	isBal(root)
	return ans
}
