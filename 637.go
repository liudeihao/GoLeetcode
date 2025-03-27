package main

func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{root}
	var ans []float64
	for len(q) > 0 {
		n := len(q)
		sum := 0
		for _ = range n {
			node := q[0]
			q = q[1:]
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(n))
	}
	return ans
}
