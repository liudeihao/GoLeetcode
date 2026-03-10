package main

type Node_559 struct {
	Val      int
	Children []*Node_559
}

func maxDepth_559(root *Node_559) int {
	if root == nil {
		return 0
	}
	var ans int
	var q []*Node_559
	q = append(q, root)
	for len(q) != 0 {
		ans++
		cq := q
		q = nil
		for _, tn := range cq {
			for _, c := range tn.Children {
				q = append(q, c)
			}
		}
	}
	return ans

}
