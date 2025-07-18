package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {

	if root == nil {
		return 0
	}
	var ans int
	var q []*Node
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
