package main

func sumNumbers(root *TreeNode) int {
	var sum int
	path := []int{root.Val}
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r.Left == nil && r.Right == nil {
			var cur int
			for _, i := range path {
				cur = cur*10 + i
			}
			sum += cur
			return
		}
		if r.Left != nil {
			path = append(path, r.Left.Val)
			dfs(r.Left)
			path = path[:len(path)-1]
		}
		if r.Right != nil {
			path = append(path, r.Right.Val)
			dfs(r.Right)
			path = path[:len(path)-1]

		}
	}
	dfs(root)
	return sum
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
