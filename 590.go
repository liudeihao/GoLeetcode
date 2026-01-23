package main

type Node_590 struct {
	Val      int
	Children []*Node_590
}

func postorder(root *Node_590) []int {
	var ans []int
	var traverse func(root *Node_590)
	traverse = func(root *Node_590) {
		if root == nil {
			return
		}
		for _, c := range root.Children {
			traverse(c)
		}
		ans = append(ans, root.Val)
	}
	traverse(root)
	return ans
}
