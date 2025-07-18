package main

func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	for len(q) > 0 {
		cq := q
		q = nil
		for _, tn := range cq {
			if tn.Left != nil {
				q = append(q, tn.Left)
			}
			if tn.Right != nil {
				q = append(q, tn.Right)
			}
		}
		if len(q) == 0 {
			return cq[0].Val
		}
	}
	return -1
}
