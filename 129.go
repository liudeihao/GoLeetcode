package main

func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := prevSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	return dfs(root, 0)
}

/*
虽然下面这个才像是一个正常的数字之和。。。

func sumNumbers(root *TreeNode) int {
	var childSumNumbers func(*TreeNode) (int, int)
	childSumNumbers = func(r *TreeNode) (sum, depth int) {
		if r == nil {
			return 0, 0
		}
		ls, ld := childSumNumbers(r.Left)
		rs, rd := childSumNumbers(r.Right)
		depth = max(ld, rd) + 1
		sum = ls + rs + depth*r.Val
		return
	}
	sum, _ := childSumNumbers(root)
	return sum
}


*/
