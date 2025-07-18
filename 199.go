package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var layer [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		cq := q
		q = nil
		var tmp []int
		for _, tn := range cq {
			tmp = append(tmp, tn.Val)
			if tn.Left != nil {
				q = append(q, tn.Left)
			}
			if tn.Right != nil {
				q = append(q, tn.Right)
			}
		}
		layer = append(layer, tmp)
	}
	var ans []int
	for _, l := range layer {
		ans = append(ans, l[len(l)-1])
	}
	return ans
}
