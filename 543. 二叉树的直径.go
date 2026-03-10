package main

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		dl, dr := depth(root.Left), depth(root.Right)
		// 只是相比于计算深度多了这么一句
		ans = max(ans, dl+dr)
		return 1 + max(dl, dr)
	}
	depth(root)
	return ans
}
