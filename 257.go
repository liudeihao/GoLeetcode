package main

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var path []string
	var bt func(r *TreeNode)
	bt = func(r *TreeNode) {
		path = append(path, strconv.Itoa(r.Val))
		if r.Left == nil && r.Right == nil {
			ans = append(ans, strings.Join(path, "->"))
		}
		if r.Left != nil {
			bt(r.Left)
		}
		if r.Right != nil {
			bt(r.Right)
		}
		path = path[:len(path)-1]
	}
	bt(root)
	return ans
}
