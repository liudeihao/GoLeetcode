package main

type Node_429 struct {
	Val      int
	Children []*Node_429
}

func levelOrder_429(root *Node_429) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	var q []*Node_429
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
