package main

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		cq := q
		q = nil
		tmp := math.MinInt
		for _, tn := range cq {
			tmp = max(tn.Val, tmp)
			if tn.Left != nil {
				q = append(q, tn.Left)
			}
			if tn.Right != nil {
				q = append(q, tn.Right)
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
