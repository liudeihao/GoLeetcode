package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	var q []*Node
	q = append(q, root)
	for len(q) != 0 {
		cq := q
		q = nil
		var tmp []int
		for _, tn := range cq {
			tmp = append(tmp, tn.Val)
			for _, c := range tn.Children {
				q = append(q, c)
			}

		}
		ans = append(ans, tmp)
	}
	return ans
}
