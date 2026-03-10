package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		cq := q
		q = nil
		cur := make([]int, len(cq))
		for i, v := range cq {
			cur[i] = v.Val
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		ans = append(ans, cur)
	}
	return ans
}
