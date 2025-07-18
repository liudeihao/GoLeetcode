package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var ans []int
	var traverse func(root *Node)
	traverse = func(root *Node) {
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
