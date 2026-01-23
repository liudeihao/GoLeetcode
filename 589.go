package main

type Node_589 struct {
	Val      int
	Children []*Node_589
}

func preorder(root *Node_589) []int {
	var ans []int
	var traverse func(root *Node_589)
	traverse = func(root *Node_589) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		for _, c := range root.Children {
			traverse(c)
		}
	}
	traverse(root)
	return ans
}
