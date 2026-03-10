package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	var ans []int

	for len(q) != 0 {
		cq := q
		q = nil
		ans = append(ans, cq[len(cq)-1].Val)
		for _, v := range cq {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return ans
}
