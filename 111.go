package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		ans++
		cq := q
		q = nil
		for _, tn := range cq {
			if tn.Left == nil && tn.Right == nil {
				return ans
			}
			if tn.Left != nil {
				q = append(q, tn.Left)
			}
			if tn.Right != nil {
				q = append(q, tn.Right)
			}
		}
	}
	return ans
}
