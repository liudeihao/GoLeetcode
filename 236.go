package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	// 二叉树后根遍历，回溯的过程本身就是从底到上。
	/*
		在递归函数有返回值的情况下：
		1) 如果要搜索一条边，递归函数返回值不为nil的时候，立刻返回，
		2) 如果搜索整个树，直接用一个变量left、right接住返回值，
			这个left、right后序还有逻辑处理的需要，也就是后序遍历中处理中间节点的逻辑（也是回溯）。
	*/
	l, r := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	} else {
		return r
	}
}
