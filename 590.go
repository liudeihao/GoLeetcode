package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var ans []int
	var traverse func(root *Node)
	traverse = func(root *Node) {
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
